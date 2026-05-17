import z from 'zod'

const VideoSchema = z.object({
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
export const VideoListSchema = z.array(VideoSchema)
