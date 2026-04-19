import { apiFetch } from "$lib/api";

  export async function uploadVideo(file: File): Promise<{ id: string }> {
    const formData = new FormData();
    formData.append('file', file);

    const response = await apiFetch('/api/videos', fetch, {                                                
        method: 'POST',
        body: formData                                                                                     
    });         

    return await response.json();
}