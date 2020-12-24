module.exports = {
  title: 'hades',
  tagline: 'The tagline of my site',
  url: 'https://everettraven.github.io',
  baseUrl: '/hades/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/hades-logo-16.png',
  organizationName: 'everettraven', // Usually your GitHub org/user name.
  projectName: 'hades', // Usually your repo name.
  themeConfig: {
    navbar: {
      title: 'hades',
      logo: {
        alt: 'My Site Logo',
        src: 'img/hades-logo-32.png',
      },
      items: [
        {
          href: 'https://github.com/everettraven/hades',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Introduction to hades',
              to: '/',
            },
            {
              label: 'Contributing to hades',
              to: '/contrib',
            },
            {
              label: 'Installing hades',
              to: '/getting_started/install',
            },
            {
              label: 'Simple Command Guide',
              to: '/guides/simple_command',
            },
            {
              label: 'Host File Guide',
              to: '/guides/host_file',
            },
            {
              label: 'Multiple Tests Guide',
              to: '/guides/multiple_tests',
            },
            {
              label: 'command',
              to: '/resources/command',
            },
            {
              label: 'os',
              to: '/resources/os',
            },
          ],
        },
        {
          title: 'More',
          items: [
            {
              label: 'GitHub',
              href: 'https://github.com/everettraven/hades',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} hades. Built with Docusaurus.`,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl:
            'https://github.com/everettraven/hades/edit/main/docs/',
            routeBasePath: '/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
};
