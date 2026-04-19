import { fetchVideoList } from '$lib/features/video/api';
import type { PageLoad } from './$types'; 

export const load: PageLoad = async ({ fetch }) => {
    const videos = await fetchVideoList(fetch);

    return { videos };
}