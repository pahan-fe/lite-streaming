import type { PageLoad } from './$types'; 
import { apiFetch } from '$lib/api';

interface VideoDto {
    id: string
    originalFilename: string
    contentType: string
    size: number
    status: string
    s3RawKey: string
    s3HLSKey: string
    createdAt: string
    updatedAt: string
}

interface Video {
    id: string
    title: string
    size: number
    status: string
    updatedAt: string
}

const mapVideoDtoToVideo = (dto: VideoDto): Video => ({
    id: dto.id,
    title: dto.originalFilename,
    size: dto.size,
    status: dto.status,
    updatedAt: new Date(dto.updatedAt).toLocaleString()
})

export const load: PageLoad = async ({ fetch }) => {
    try {
        const response = await apiFetch('/api/videos', fetch);
        const result = await response.json();

        return { videos: result.map(mapVideoDtoToVideo) };
    } catch (error) {
        console.error('Failed to load video list:', error);
    }
}