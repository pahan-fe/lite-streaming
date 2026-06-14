import z from 'zod'

export const VideoSchema = z.object({
  id: z.string(),
  originalFilename: z.string(),
  contentType: z.string(),
  size: z.number(),
  status: z.string(),
  s3RawKey: z.string(),
  s3HLSKey: z.string(),
  createdAt: z.coerce.date(),
  updatedAt: z.coerce.date(),
})
export type Video = z.infer<typeof VideoSchema>

export const VideoListResponseSchema = z.object({
  items: z.array(VideoSchema),
  nextCursor: z.string().nullable(),
})
export type VideoListResponse = z.infer<typeof VideoListResponseSchema>
