// file: sdks/javascript/src/index.ts
// version: 1.0.0
// guid: 550e8400-e29b-41d4-a716-446655440012

/**
 * Subtitle Manager JavaScript/TypeScript SDK
 *
 * A comprehensive SDK for the Subtitle Manager API with full TypeScript support,
 * automatic retry, error handling, and modern async/await patterns.
 */

export { SubtitleManagerClient } from './client';
export * from './types';
export * from './errors';
export * from './models';

// Default export for CommonJS compatibility
import { SubtitleManagerClient } from './client';
export default SubtitleManagerClient;
