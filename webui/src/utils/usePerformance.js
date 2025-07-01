// file: webui/src/utils/usePerformance.js
// version: 1.0.0
// guid: 5f4e3d2c-1f0e-6f5e-9a8b-2f1e0f9e8f7e

import { useState, useEffect, useCallback } from 'react';

/**
 * Custom hook for monitoring frontend performance metrics.
 * 
 * This hook provides:
 * - Page load performance tracking
 * - Component render time monitoring
 * - Memory usage tracking
 * - Network request performance
 * - User interaction metrics
 * 
 * @param {Object} options - Configuration options
 * @param {boolean} [options.enabled=true] - Whether to collect metrics
 * @param {number} [options.updateInterval=30000] - How often to update metrics (ms)
 * @param {string} [options.componentName] - Name of component for tracking
 * 
 * @returns {Object} Performance metrics and utilities
 */
export const usePerformance = (options = {}) => {
  const {
    enabled = true,
    updateInterval = 30000,
    componentName = 'Unknown',
  } = options;

  const [metrics, setMetrics] = useState({
    loadTime: 0,
    renderTime: 0,
    memoryUsage: 0,
    fps: 0,
    networkRequests: 0,
    cacheHitRatio: 0,
  });

  const [isCollecting, setIsCollecting] = useState(false);

  // Track component render time
  const trackRenderTime = useCallback((startTime) => {
    if (!enabled) return;
    
    const endTime = performance.now();
    const renderTime = endTime - startTime;
    
    setMetrics(prev => ({
      ...prev,
      renderTime: Math.round(renderTime * 100) / 100,
    }));
  }, [enabled]);

  // Start render timing
  const startRenderTiming = useCallback(() => {
    if (!enabled) return null;
    return performance.now();
  }, [enabled]);

  // Track user interaction performance
  const trackInteraction = useCallback((interactionName, duration) => {
    if (!enabled) return;
    
    // Could send to analytics service
    console.debug(`Interaction ${interactionName} took ${duration}ms`);
  }, [enabled]);

  // Get current performance metrics
  const getPerformanceMetrics = useCallback(() => {
    if (!enabled || typeof performance === 'undefined') {
      return {
        loadTime: 0,
        memoryUsage: 0,
        timing: {},
      };
    }

    // Get navigation timing
    const navigation = performance.getEntriesByType('navigation')[0];
    const loadTime = navigation ? navigation.loadEventEnd - navigation.fetchStart : 0;

    // Get memory usage (if available)
    let memoryUsage = 0;
    if (performance.memory) {
      memoryUsage = Math.round(performance.memory.usedJSHeapSize / 1024 / 1024 * 100) / 100;
    }

    // Get resource timing for network requests
    const resources = performance.getEntriesByType('resource');
    const networkRequests = resources.length;

    // Calculate cache hit ratio from resource timing
    const cachedResources = resources.filter(resource => 
      resource.transferSize === 0 && resource.decodedBodySize > 0
    );
    const cacheHitRatio = networkRequests > 0 
      ? Math.round((cachedResources.length / networkRequests) * 100) 
      : 0;

    return {
      loadTime: Math.round(loadTime),
      memoryUsage,
      networkRequests,
      cacheHitRatio,
      timing: {
        dns: navigation ? navigation.domainLookupEnd - navigation.domainLookupStart : 0,
        connect: navigation ? navigation.connectEnd - navigation.connectStart : 0,
        request: navigation ? navigation.responseStart - navigation.requestStart : 0,
        response: navigation ? navigation.responseEnd - navigation.responseStart : 0,
        dom: navigation ? navigation.domContentLoadedEventEnd - navigation.domContentLoadedEventStart : 0,
      },
    };
  }, [enabled]);

  // Monitor FPS (frames per second)
  const monitorFPS = useCallback(() => {
    if (!enabled || typeof requestAnimationFrame === 'undefined') return;

    let frames = 0;
    let startTime = performance.now();

    const countFrame = () => {
      frames++;
      const currentTime = performance.now();
      
      if (currentTime - startTime >= 1000) {
        setMetrics(prev => ({
          ...prev,
          fps: frames,
        }));
        frames = 0;
        startTime = currentTime;
      }
      
      if (isCollecting) {
        requestAnimationFrame(countFrame);
      }
    };

    if (isCollecting) {
      requestAnimationFrame(countFrame);
    }
  }, [enabled, isCollecting]);

  // Start collecting performance metrics
  const startCollecting = useCallback(() => {
    if (!enabled) return;
    
    setIsCollecting(true);
    
    // Get initial metrics
    const initialMetrics = getPerformanceMetrics();
    setMetrics(prev => ({
      ...prev,
      ...initialMetrics,
    }));
  }, [enabled, getPerformanceMetrics]);

  // Stop collecting performance metrics
  const stopCollecting = useCallback(() => {
    setIsCollecting(false);
  }, []);

  // Update metrics periodically
  useEffect(() => {
    if (!enabled || !isCollecting) return;

    const interval = setInterval(() => {
      const currentMetrics = getPerformanceMetrics();
      setMetrics(prev => ({
        ...prev,
        ...currentMetrics,
      }));
    }, updateInterval);

    return () => clearInterval(interval);
  }, [enabled, isCollecting, updateInterval, getPerformanceMetrics]);

  // Start FPS monitoring when collecting
  useEffect(() => {
    monitorFPS();
  }, [monitorFPS]);

  // Auto-start collecting on mount
  useEffect(() => {
    if (enabled) {
      startCollecting();
    }
    
    return () => {
      stopCollecting();
    };
  }, [enabled, startCollecting, stopCollecting]);

  // Report performance to backend
  const reportPerformance = useCallback(async (additionalData = {}) => {
    if (!enabled) return;

    const performanceData = {
      componentName,
      timestamp: new Date().toISOString(),
      metrics: getPerformanceMetrics(),
      userAgent: navigator.userAgent,
      url: window.location.href,
      ...additionalData,
    };

    try {
      // Send to backend performance endpoint
      await fetch('/api/performance/frontend', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(performanceData),
      });
    } catch (error) {
      console.warn('Failed to report performance metrics:', error);
    }
  }, [enabled, componentName, getPerformanceMetrics]);

  return {
    metrics,
    isCollecting,
    startCollecting,
    stopCollecting,
    trackRenderTime,
    startRenderTiming,
    trackInteraction,
    reportPerformance,
    getPerformanceMetrics,
  };
};

/**
 * Hook for tracking component-specific performance metrics.
 * 
 * @param {string} componentName - Name of the component
 * @param {Object} options - Additional options
 * 
 * @returns {Object} Component performance utilities
 */
export const useComponentPerformance = (componentName, options = {}) => {
  const performance = usePerformance({
    ...options,
    componentName,
  });

  // Track component mount time
  useEffect(() => {
    const startTime = performance.startRenderTiming();
    
    // Track mount completion in next tick
    const timeoutId = setTimeout(() => {
      if (startTime) {
        performance.trackRenderTime(startTime);
      }
    }, 0);

    return () => clearTimeout(timeoutId);
  }, [performance]);

  return performance;
};

/**
 * Higher-order component for automatic performance tracking.
 * 
 * @param {React.Component} WrappedComponent - Component to wrap
 * @param {string} componentName - Name for tracking
 * 
 * @returns {React.Component} Enhanced component with performance tracking
 */
export const withPerformanceTracking = (WrappedComponent, componentName) => {
  return function PerformanceTrackedComponent(props) {
    const performance = useComponentPerformance(componentName);
    
    return <WrappedComponent {...props} performance={performance} />;
  };
};

/**
 * Utility function to measure async operation performance.
 * 
 * @param {Function} asyncOperation - Async function to measure
 * @param {string} operationName - Name of the operation
 * 
 * @returns {Promise} Result of the async operation
 */
export const measureAsync = async (asyncOperation, operationName) => {
  const startTime = performance.now();
  
  try {
    const result = await asyncOperation();
    const duration = performance.now() - startTime;
    
    console.debug(`${operationName} completed in ${Math.round(duration)}ms`);
    
    return result;
  } catch (error) {
    const duration = performance.now() - startTime;
    console.warn(`${operationName} failed after ${Math.round(duration)}ms:`, error);
    throw error;
  }
};

export default usePerformance;