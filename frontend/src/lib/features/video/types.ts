export interface VideoDto {
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

export interface Video {
    id: string
    title: string
    size: number
    status: string
    updatedAt: string
}