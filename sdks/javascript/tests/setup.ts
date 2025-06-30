// file: sdks/javascript/tests/setup.ts
// version: 1.0.0
// guid: 550e8400-e29b-41d4-a716-446655440018

/**
 * Test setup configuration
 */

// Mock environment variables
process.env.SUBTITLE_MANAGER_API_KEY = 'test-api-key';

// Mock FormData for Node.js environment
if (typeof FormData === 'undefined') {
  global.FormData = class FormData {
    private data: Map<string, any> = new Map();

    append(key: string, value: any, filename?: string): void {
      this.data.set(key, { value, filename });
    }

    get(key: string): any {
      return this.data.get(key)?.value;
    }

    has(key: string): boolean {
      return this.data.has(key);
    }

    delete(key: string): void {
      this.data.delete(key);
    }

    entries(): Iterator<[string, any]> {
      return this.data.entries();
    }
  } as any;
}

// Mock File and Blob for Node.js environment
if (typeof File === 'undefined') {
  global.File = class File {
    constructor(
      public content: any[],
      public name: string,
      public options: any = {}
    ) {}
  } as any;
}

if (typeof Blob === 'undefined') {
  global.Blob = class Blob {
    constructor(public content: any[], public options: any = {}) {}
  } as any;
}