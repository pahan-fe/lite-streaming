export const formatDate = (date: Date, month: 'short' | 'long' = 'short') =>
  new Intl.DateTimeFormat('en-US', {
    month,
    day: 'numeric',
    year: 'numeric',
    timeZone: 'UTC',
  }).format(date)
