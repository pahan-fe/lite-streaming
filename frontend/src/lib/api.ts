import { browser } from '$app/environment';

export async function apiFetch(path: string, fetchFn: typeof fetch = fetch, init?: RequestInit): Promise<Response> {
    let response: Response;

    if (browser) {
       response = await fetchFn(path, init);
    } else {
        const { env } = await import('$env/dynamic/private');
        const url = `${env.INTERNAL_API_URL ?? 'http://localhost:8080'}${path}`;
    
        response = await globalThis.fetch(url, init);
    }

    if (!response.ok) {
        const message = await response.text();
        throw new Error( `API request failed: ${message.trim() || response.statusText} (status ${response.status})`);
    }

    return response;
}