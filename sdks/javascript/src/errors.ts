// file: sdks/javascript/src/errors.ts
// version: 1.0.0
// guid: 550e8400-e29b-41d4-a716-446655440014

/**
 * Error classes for the Subtitle Manager SDK
 */

export class SubtitleManagerError extends Error {
  public readonly statusCode?: number;
  public readonly errorCode?: string;
  public readonly details?: Record<string, any>;

  constructor(
    message: string,
    statusCode?: number,
    errorCode?: string,
    details?: Record<string, any>
  ) {
    super(message);
    this.name = 'SubtitleManagerError';
    this.statusCode = statusCode;
    this.errorCode = errorCode;
    this.details = details;

    // Maintain proper stack trace for where our error was thrown (only available on V8)
    if (Error.captureStackTrace) {
      Error.captureStackTrace(this, SubtitleManagerError);
    }
  }
}

export class AuthenticationError extends SubtitleManagerError {
  constructor(message: string, statusCode?: number, errorCode?: string, details?: Record<string, any>) {
    super(message, statusCode, errorCode, details);
    this.name = 'AuthenticationError';
  }
}

export class AuthorizationError extends SubtitleManagerError {
  constructor(message: string, statusCode?: number, errorCode?: string, details?: Record<string, any>) {
    super(message, statusCode, errorCode, details);
    this.name = 'AuthorizationError';
  }
}

export class NotFoundError extends SubtitleManagerError {
  constructor(message: string, statusCode?: number, errorCode?: string, details?: Record<string, any>) {
    super(message, statusCode, errorCode, details);
    this.name = 'NotFoundError';
  }
}

export class RateLimitError extends SubtitleManagerError {
  public readonly retryAfter?: number;

  constructor(
    message: string,
    retryAfter?: number,
    statusCode?: number,
    errorCode?: string,
    details?: Record<string, any>
  ) {
    super(message, statusCode, errorCode, details);
    this.name = 'RateLimitError';
    this.retryAfter = retryAfter;
  }
}

export class ValidationError extends SubtitleManagerError {
  constructor(message: string, statusCode?: number, errorCode?: string, details?: Record<string, any>) {
    super(message, statusCode, errorCode, details);
    this.name = 'ValidationError';
  }
}

export class NetworkError extends SubtitleManagerError {
  constructor(message: string, originalError?: Error) {
    super(message);
    this.name = 'NetworkError';
    
    if (originalError) {
      this.details = { originalError: originalError.message };
      this.stack = originalError.stack;
    }
  }
}

export class TimeoutError extends SubtitleManagerError {
  constructor(message: string) {
    super(message);
    this.name = 'TimeoutError';
  }
}

export class APIError extends SubtitleManagerError {
  constructor(message: string, statusCode?: number, errorCode?: string, details?: Record<string, any>) {
    super(message, statusCode, errorCode, details);
    this.name = 'APIError';
  }
}