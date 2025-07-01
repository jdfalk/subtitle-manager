// file: webui/src/utils/LazyImage.jsx
// version: 1.0.0
// guid: 4f3e2d1c-0f9e-5f4e-8a7b-1f0e9f8e7f6e

import React, { useState, useRef, useEffect } from 'react';

/**
 * LazyImage component provides lazy loading functionality for images to improve performance.
 * 
 * Features:
 * - Loads images only when they enter the viewport
 * - Shows placeholder while loading
 * - Handles loading and error states
 * - Optimizes memory usage and initial page load time
 * - Supports responsive images with srcSet
 * 
 * @param {Object} props - Component props
 * @param {string} props.src - Image source URL
 * @param {string} props.alt - Alt text for accessibility
 * @param {string} [props.placeholder] - Placeholder image while loading
 * @param {string} [props.className] - CSS classes to apply
 * @param {Object} [props.style] - Inline styles
 * @param {string} [props.srcSet] - Responsive image sources
 * @param {string} [props.sizes] - Image sizes for responsive loading
 * @param {Function} [props.onLoad] - Callback when image loads
 * @param {Function} [props.onError] - Callback when image fails to load
 * @param {number} [props.threshold=0.1] - Intersection threshold (0-1)
 * @param {string} [props.rootMargin='50px'] - Root margin for intersection
 */
const LazyImage = ({
  src,
  alt,
  placeholder,
  className = '',
  style = {},
  srcSet,
  sizes,
  onLoad,
  onError,
  threshold = 0.1,
  rootMargin = '50px',
  ...props
}) => {
  const [isLoaded, setIsLoaded] = useState(false);
  const [isInView, setIsInView] = useState(false);
  const [hasError, setHasError] = useState(false);
  const imgRef = useRef();

  // Intersection Observer to detect when image enters viewport
  useEffect(() => {
    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          setIsInView(true);
          observer.disconnect();
        }
      },
      {
        threshold,
        rootMargin,
      }
    );

    if (imgRef.current) {
      observer.observe(imgRef.current);
    }

    return () => observer.disconnect();
  }, [threshold, rootMargin]);

  // Handle image load
  const handleLoad = (event) => {
    setIsLoaded(true);
    onLoad?.(event);
  };

  // Handle image error
  const handleError = (event) => {
    setHasError(true);
    onError?.(event);
  };

  // Render placeholder while not in view or loading
  if (!isInView) {
    return (
      <div
        ref={imgRef}
        className={`lazy-image-placeholder ${className}`}
        style={{
          backgroundColor: '#f0f0f0',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          minHeight: '100px',
          ...style,
        }}
        {...props}
      >
        {placeholder ? (
          <img src={placeholder} alt={alt} style={{ maxWidth: '100%', maxHeight: '100%' }} />
        ) : (
          <div style={{ color: '#999', fontSize: '14px' }}>Loading...</div>
        )}
      </div>
    );
  }

  // Render error state
  if (hasError) {
    return (
      <div
        className={`lazy-image-error ${className}`}
        style={{
          backgroundColor: '#f8f8f8',
          border: '1px solid #ddd',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          minHeight: '100px',
          color: '#666',
          fontSize: '14px',
          ...style,
        }}
        {...props}
      >
        Failed to load image
      </div>
    );
  }

  // Render actual image when in view
  return (
    <img
      ref={imgRef}
      src={src}
      srcSet={srcSet}
      sizes={sizes}
      alt={alt}
      className={`lazy-image ${isLoaded ? 'loaded' : 'loading'} ${className}`}
      style={{
        opacity: isLoaded ? 1 : 0.5,
        transition: 'opacity 0.3s ease',
        ...style,
      }}
      onLoad={handleLoad}
      onError={handleError}
      {...props}
    />
  );
};

export default LazyImage;