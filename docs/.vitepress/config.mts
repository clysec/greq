import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "GREQ - Simple Go Requests",
  description: "A simple request library for Go",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Getting Started', link: '/getting-started' },
      { text: 'API Documentation', link: '/api-docs' }
    ],

    sidebar: [
      {
        text: 'Introduction',
        items: [
          { text: 'Getting Started', link: '/getting-started' },
          { text: 'Query and Headers', link: '/query-and-headers' }
        ]
      },
      {
        text: 'Request Types',
        items: [
          { text: 'GET', link: '/get-request' },
          { text: 'POST', link: '/post-request' },
          { text: 'PUT', link: '/put-request' },
          { text: 'PATCH', link: '/patch-request' },
          { text: 'DELETE', link: '/delete-request' },
        ]
      },
      {
        text: 'Body Types',
        items: [
          { text: 'JSON/XML', link: '/body-marshal'},
          { text: 'URL-Encoded Form', link: '/body-form' },
          { text: 'Multipart Form', link: '/body-multipart' },
          { text: 'Raw String/Bytes', link: '/body-raw' },
          { text: 'SOAP Request', link: '/body-soap' },
          { text: 'GraphQL Request', link: '/body-graphql' },
        ],
      },
      {
        text: 'Response Types',
        items: [
          { text: 'JSON/XML', link: '/response-marshal' },
          { text: 'Raw String/Bytes', link: '/response-raw' },
          { text: 'SOAP Response', link: '/response-soap' },
        ]
      },
      {
        text: 'Authentication Helpers',
        items: [
          { text: 'Basic', link: '/auth-basic' },
          { text: 'AWS', link: '/auth-aws' },
          { text: 'JWT', link: '/auth-jwt' },
          { text: 'NTLM', link: '/auth-ntlm' },
          { text: 'Oauth2', link: '/auth-oauth2' },
          { text: 'Header', link: '/auth-header' },
          { text: 'Bearer Token', link: '/auth-bearer' },
          { text: 'mTLS/Cert Auth', link: '/auth-cert' },
          { text: 'Custom Modules', link: '/auth-custom' },
        ]
      },
      
      
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/clysec/greq' }
    ]
  }
})
