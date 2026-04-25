import { describe, it, expect } from 'vitest';
import { mapVideoDtoToVideo } from './mappers';
import type { VideoDto } from './types';

const baseDto: VideoDto = {
    id: 'abc-123',
    originalFilename: 'beach.mp4',
    contentType: 'video/mp4',
    size: 1024,
    status: 'ready',
    s3RawKey: 'videos/abc-123/original/beach.mp4',
    s3HLSKey: 'videos/abc-123/hls',
    createdAt: '2026-04-25T10:00:00Z',
    updatedAt: '2026-04-25T10:30:00Z'
};

describe('mapVideoDtoToVideo', () => {
    it('maps id and size as-is', () => {
        const result = mapVideoDtoToVideo(baseDto);
        expect(result.id).toBe('abc-123');
        expect(result.size).toBe(1024);
    });

    it('maps originalFilename to title', () => {
        const result = mapVideoDtoToVideo(baseDto);
        expect(result.title).toBe('beach.mp4');
    });

    it('maps status as-is', () => {
        expect(mapVideoDtoToVideo({ ...baseDto, status: 'ready' }).status).toBe('ready');
        expect(mapVideoDtoToVideo({ ...baseDto, status: 'processing' }).status).toBe('processing');
        expect(mapVideoDtoToVideo({ ...baseDto, status: 'failed' }).status).toBe('failed');
    });

    it('formats updatedAt via toLocaleString', () => {
        const expected = new Date(baseDto.updatedAt).toLocaleString();
        expect(mapVideoDtoToVideo(baseDto).updatedAt).toBe(expected);
    });

    it('does not leak DTO-only fields into Video', () => {
        const result = mapVideoDtoToVideo(baseDto);
        expect(result).not.toHaveProperty('s3RawKey');
        expect(result).not.toHaveProperty('s3HLSKey');
        expect(result).not.toHaveProperty('contentType');
        expect(result).not.toHaveProperty('createdAt');
        expect(result).not.toHaveProperty('originalFilename');
    });
});
