import { fetchVideoById } from "$lib/features/video/api";
 import type { PageLoad } from "./$types"; 

export const load: PageLoad = async ({ fetch, params, depends }) => {
    depends(`video:${params.id}`);

    const video = await fetchVideoById(params.id, fetch);

    return { video };
}