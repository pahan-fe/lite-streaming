<template>
  <div>
    <header class="border-b">
      <div class="container mx-auto flex items-center justify-between py-4">
        <NuxtLink to="/" class="text-xl font-bold">lite-streaming</NuxtLink>
        <nav class="flex gap-2">
          <Button as-child variant="ghost">
            <NuxtLink to="/" exact-active-class="text-foreground underline">Library</NuxtLink>
          </Button>
          <Button as-child variant="ghost">
            <NuxtLink to="/upload" active-class="text-foreground underline">Upload</NuxtLink>
          </Button>
        </nav>
      </div>
    </header>
    <main class="container mx-auto py-8">
      <slot />
    </main>
  </div>
</template>
