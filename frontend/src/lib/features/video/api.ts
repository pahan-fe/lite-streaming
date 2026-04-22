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

export const fetchVideoById = async (id: string, fetchFn: typeof fetch): Promise<Video> => {
  const response = await apiFetch(`/api/videos/${id}`, fetchFn);
  const result = await response.json();

  return mapVideoDtoToVideo(result);
}

export const deleteVideoById = async (id: string, fetchFn: typeof fetch): Promise<void> => {
  const response = await apiFetch(`/api/videos/${id}`, fetchFn, { method: 'DELETE' });
  if (!response.ok) {
    throw new Error(`Delete failed: ${response.status}`);
  }
}