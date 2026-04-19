import type { VideoDto, Video } from "./types";

export const mapVideoDtoToVideo = (dto: VideoDto): Video => ({
    id: dto.id,
    title: dto.originalFilename,
    size: dto.size,
    status: dto.status,
    updatedAt: new Date(dto.updatedAt).toLocaleString()
})