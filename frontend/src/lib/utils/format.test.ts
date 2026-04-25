import { describe, it, expect } from 'vitest';
import { formatSize } from './format';

describe('formatSize', () => {
    it('formats bytes under 1KB as B', () => {
        expect(formatSize(0)).toBe('0 B');
        expect(formatSize(1)).toBe('1 B');
        expect(formatSize(512)).toBe('512 B');
        expect(formatSize(1023)).toBe('1023 B');
    });

    it('formats KB with one decimal', () => {
        expect(formatSize(1024)).toBe('1.0 KB');
        expect(formatSize(1536)).toBe('1.5 KB');
        expect(formatSize(1024 * 1023)).toBe('1023.0 KB');
    });

    it('formats MB with one decimal', () => {
        expect(formatSize(1024 * 1024)).toBe('1.0 MB');
        expect(formatSize(1024 * 1024 * 5.5)).toBe('5.5 MB');
        expect(formatSize(1024 * 1024 * 1023)).toBe('1023.0 MB');
    });

    it('formats GB with two decimals', () => {
        expect(formatSize(1024 ** 3)).toBe('1.00 GB');
        expect(formatSize(1024 ** 3 * 2.5)).toBe('2.50 GB');
        expect(formatSize(1024 ** 3 * 12.345)).toBe('12.35 GB');
    });

    it('crosses unit boundaries correctly', () => {
        expect(formatSize(1024 * 1024 - 1)).toBe('1024.0 KB');
        expect(formatSize(1024 ** 3 - 1)).toBe('1024.0 MB');
    });
});
