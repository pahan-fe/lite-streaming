import { apiFetch } from "$lib/api";
import { mapVideoDtoToVideo } from "./mappers";
import type { Video } from "./types";

export const uploadVideo = async (file: File): Promise<{ id: string }> => {
  const formData = new FormData();
  formData.append('file', file);

  const response = await apiFetch('/api/videos', fetch, {                                                
      method: 'POST',
      body: formData
  });         

  return await response.json();
}

export const fetchVideoList = async (fetchFn: typeof fetch): Promise<Video[]> => {
  const response = await apiFetch('/api/videos', fetchFn);
  const result = await response.json();

  return result.map(mapVideoDtoToVideo)
}