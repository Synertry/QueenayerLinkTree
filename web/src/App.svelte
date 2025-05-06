<script lang='ts'>

  const profile: {
    readonly avatar: {
      readonly src: string,
      readonly link: string,
      readonly deco: {
        readonly src: string,
        readonly link: string
      }
    },
    readonly name: string,
    readonly bio: string
  } = {
    // Image links
    avatar: {
      src: new URL('assets/images/queenayer.webp', import.meta.url).href,
      link: '/profilepic-link',

      deco: {
        src: new URL('assets/images/discord_avatardeco_angel.webp', import.meta.url).href,
        link: '/avatardeco-link',
      },
    },

    // Info
    name: 'Queen Ayer',
    bio: `
      ✨ Battling Honkai, exploring Miraland, being pretty in the Rift
      ~ Miss Pink Elf♪!
    `,
  }

  // Social media links with actual URLs
  import Icon from '@iconify/svelte';

  const links: ReadonlyArray<{
    readonly name: string,
    readonly url: string,
    readonly icon: string
  }> = [
    {name: 'Discord', url: '/discord', icon: 'simple-icons:discord'},
    {name: 'Steam', url: '/steam', icon: 'simple-icons:steam'},
    {name: 'Twitter', url: '/twitter', icon: 'simple-icons:x'},
    {name: 'Reddit', url: '/reddit', icon: 'simple-icons:reddit'},
    {name: 'Youtube', url: '/youtube', icon: 'simple-icons:youtube'},
    {name: 'Twitch', url: '/twitch', icon: 'simple-icons:twitch'},
  ];

  // Conillo mascot link randomizer
  const conillos: ReadonlyArray<string> = [
    new URL('assets/images/conillo/burrito_conillo.webp', import.meta.url).href,
    new URL('assets/images/conillo/fun_conillo.webp', import.meta.url).href,
    new URL('assets/images/conillo/surprise_gift_conillo.webp', import.meta.url).href,
    new URL('assets/images/conillo/wave_hello_conillo.webp', import.meta.url).href,
    new URL('assets/images/conillo/yoga_bunny_conillo.webp', import.meta.url).href,
  ]

  const conillo: {
    src: string,
    link: string
  } = {
    src: conillos[Math.floor(Math.random() * conillos.length)],
    link: '/conillo-link',
  }
</script>

<main>
  <div class="petals-container">
    {#each Array(20) as _, i}
      <div class="petal" style="--i: {i}"></div>
    {/each}
  </div>

  <div class="profile">
    <div class="avatar-container">
      <!-- Profile picture without link -->
      <img src={profile.avatar.src} alt="{profile.name} Avatar" class="avatar" />
      <!-- Invisible proxy circle for clickable area -->
      <a href={profile.avatar.link} target="_blank" rel="noopener noreferrer" class="avatar-proxy-link" aria-label="View profile picture">
        <div class="avatar-proxy"></div>
      </a>
      <!-- Avatar decoration in-between -->
      <a href={profile.avatar.deco.link} target="_blank" rel="noopener noreferrer" class="avatar-deco-link" aria-label="View avatar decoration">
        <img src={profile.avatar.deco.src} alt="Avatar Decoration" class="avatar-deco" />
      </a>
    </div>
    <h1 class="username">{profile.name}</h1>
    <p class="bio">{profile.bio}</p>
  </div>

  <div class="links">
    {#each links as link}
      <a href={link.url} target="_blank" rel="noopener noreferrer" class="link-button" aria-label="Link to {link.name} profile">
        <div class="icon-container">
          <Icon icon={link.icon} width="1.5em" height="1.5em" />
        </div>
        <span>{link.name}</span>
      </a>
    {/each}
  </div>

  <div class="mascot-container">
    <a href={conillo.link} target="_blank" rel="noopener noreferrer" class="mascot-link" aria-label="View Conillo mascot">
      <img src={conillo.src} alt="Conillo mascot" class="mascot" />
    </a>
  </div>

  <footer class="footer">
    Made with ❤️ by <a href="/synertry" target="_blank" rel="noopener noreferrer">Synertry</a>
  </footer>
</main>