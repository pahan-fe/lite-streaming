import { fetchVideoList } from '$lib/features/video/api';
import { pageSize } from '$lib/features/video/const';
import type { PageLoad } from './$types'; 

export const load: PageLoad = async ({ fetch }) => {
    const videos = await fetchVideoList(fetch, 1, pageSize);

    return { videos };
}