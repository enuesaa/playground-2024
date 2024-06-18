import a from 'reg-cli'

a({
  actualDir: "path/to/actual/directory",
  expectedDir: "path/to/expected/directory",
  diffDir: "path/to/save/diff/images",
  report: "path/to/save/report/file", // Optional
  junitReport: "path/to/save/junit/report", // Optional
  json: "path/to/save/json/results", // Optional
  update: true, // Optional boolean flag
  extendedErrors: false, // Optional boolean flag
  urlPrefix: "https://example.com/images/", // Optional URL prefix
  matchingThreshold: 0.8, // Optional number (0-1)
  threshold: 0.8, // Alias to thresholdRate; Optional number (0-1)
  thresholdRate: 0.8, // Optional number (0-1)
  thresholdPixel: 10, // Optional number
  concurrency: 4, // Optional number
  enableAntialias: true, // Optional boolean flag
  enableClientAdditionalDetection: false // Optional boolean flag
})
