import { browser } from '$app/environment';

export async function apiFetch(path: string, fetchFn: typeof fetch = fetch, init?: RequestInit): Promise<Response> {
    let url = path;

    if (!browser) {
        const { env } = await import('$env/dynamic/private');
        url = `${env.INTERNAL_API_URL ?? 'http://localhost:8080'}${path}`;
    }

    return fetchFn(url, init);                                                     
}