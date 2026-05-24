<template>
  <div class="min-h-screen">
    <header class="sticky top-0 z-40 border-b border-border/60 bg-background/70 backdrop-blur-xl">
      <div class="container mx-auto flex h-16 items-center justify-between px-6">
        <NuxtLink to="/" class="group flex items-center gap-2.5">
          <span
            class="grid size-7 place-items-center rounded-md bg-primary/15 text-primary transition-transform duration-300 group-hover:scale-110"
          >
            <svg viewBox="0 0 24 24" class="size-3.5 fill-current" aria-hidden="true">
              <path d="M8 5v14l11-7z" />
            </svg>
          </span>
          <span class="font-display text-xl font-medium tracking-tight">
            lite<span class="text-primary">·</span>streaming
          </span>
        </NuxtLink>

        <nav class="flex items-center gap-1">
          <NuxtLink
            to="/"
            exact-active-class="text-foreground after:scale-x-100"
            class="relative px-3 py-2 text-sm text-muted-foreground transition-colors after:absolute after:inset-x-3 after:bottom-1 after:h-px after:origin-left after:scale-x-0 after:bg-primary after:transition-transform after:duration-300 hover:text-foreground"
          >
            Library
          </NuxtLink>
          <NuxtLink
            to="/upload"
            active-class="text-foreground after:scale-x-100"
            class="relative px-3 py-2 text-sm text-muted-foreground transition-colors after:absolute after:inset-x-3 after:bottom-1 after:h-px after:origin-left after:scale-x-0 after:bg-primary after:transition-transform after:duration-300 hover:text-foreground"
          >
            Upload
          </NuxtLink>
        </nav>
      </div>
    </header>

    <main class="container mx-auto px-6 py-12">
      <slot />
    </main>
  </div>
</template>
