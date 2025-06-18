// file: webui/src/utils/__tests__/configSanitizer.test.js

import { describe, expect, test } from 'vitest';
import { sanitizeConfig } from '../configSanitizer.js';

describe('configSanitizer utility', () => {
  describe('sanitizeConfig', () => {
    test('masks sensitive values by default', () => {
      const config = {
        openai_api_key: 'sk-abcdef123456',
        database_password: 'secretpassword',
        normal_value: 'visible',
      };

      const sanitized = sanitizeConfig(config);

      expect(sanitized.openai_api_key).toBe('****3456');
      expect(sanitized.database_password).toBe('****word');
      expect(sanitized.normal_value).toBe('visible');
    });

    test('shows sensitive values when showSensitive is true', () => {
      const config = {
        openai_api_key: 'sk-abcdef123456',
        database_password: 'secretpassword',
        normal_value: 'visible',
      };

      const sanitized = sanitizeConfig(config, true);

      expect(sanitized.openai_api_key).toBe('sk-abcdef123456');
      expect(sanitized.database_password).toBe('secretpassword');
      expect(sanitized.normal_value).toBe('visible');
    });

    test('handles nested objects', () => {
      const config = {
        database: {
          password: 'secretpass',
          host: 'localhost',
        },
        api: {
          token: 'abc123',
        },
      };

      const sanitized = sanitizeConfig(config);

      expect(sanitized.database.password).toBe('****pass');
      expect(sanitized.database.host).toBe('localhost');
      expect(sanitized.api.token).toBe('****c123');
    });

    test('handles arrays', () => {
      const config = {
        keys: ['api_key_1', 'api_key_2'],
        passwords: ['pass1', 'pass2'],
        normal: ['value1', 'value2'],
      };

      const sanitized = sanitizeConfig(config);

      // When the array key itself is sensitive, sanitize the array values
      expect(sanitized.keys).toEqual(['****ey_1', '****ey_2']);
      expect(sanitized.passwords).toEqual(['****ass1', '****ass2']);
      expect(sanitized.normal).toEqual(['value1', 'value2']);
    });

    test('handles non-string sensitive values', () => {
      const config = {
        password: 123456,
        api_key: null,
        token: undefined,
      };

      const sanitized = sanitizeConfig(config);

      expect(sanitized.password).toBe('****');
      expect(sanitized.api_key).toBe('****');
      expect(sanitized.token).toBe('****');
    });

    test('handles null and undefined inputs', () => {
      expect(sanitizeConfig(null)).toBe(null);
      expect(sanitizeConfig(undefined)).toBe(undefined);
    });

    test('identifies various sensitive key patterns', () => {
      const config = {
        password: 'secret1',
        PASSWORD: 'secret2',
        apikey: 'secret3',
        api_key: 'secret4',
        API_KEY: 'secret5',
        token: 'secret6',
        TOKEN: 'secret7',
        secret: 'secret8',
        SECRET: 'secret9',
        oauth_token: 'secret10',
        database_password: 'secret11',
      };

      const sanitized = sanitizeConfig(config);

      Object.values(sanitized).forEach(value => {
        expect(value).toMatch(/^\*\*\*\*/);
      });
    });
  });
});
