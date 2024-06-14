/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: false, // for webcontainer
  experimental: {
    appDir: false,
  },
  async headers() {
    const securityHeaders = [
      {
        key: 'Cross-Origin-Embedder-Policy',
        value: 'require-corp'
      },
      {
        key: 'Cross-Origin-Opener-Policy',
        value: 'same-origin'
      }
    ]

    return [
      {
        source: '/:path*',
        headers: securityHeaders,
      },
    ]
  },
}

module.exports = nextConfig
