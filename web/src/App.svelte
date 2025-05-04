<script lang="ts">
  import Icon from '@iconify/svelte';

  // Image links
  const profilepic = "/profilepic";
  const profilepicLink = "/profilepic-link";
  const avatardeco = "/avatardeco";
  const avatardecoLink = "/avatardeco-link";
  const conillo = "/conillo";
  const conilloLink = "/conillo-link";

  // multiline string for the bio
  const bio = `
    ✨ Battling Honkai, exploring Miraland, being pretty in the Rift
    ~ Miss Pink Elf♪!
  `;


  // Social media links with actual URLs
  const links = [
    { name: 'Discord', url: '/discord', icon: 'simple-icons:discord' },
    { name: 'Steam', url: '/steam', icon: 'simple-icons:steam' },
    { name: 'Twitter', url: '/twitter', icon: 'simple-icons:x' },
    { name: 'Reddit', url: '/reddit', icon: 'simple-icons:reddit' },
    { name: 'Youtube', url: '/youtube', icon: 'simple-icons:youtube' },
    { name: 'Twitch', url: '/twitch', icon: 'simple-icons:twitch' },
  ];
</script>

<main>

  <!-- Petal Animation Container -->
  <div class="petals-container">
    {#each Array(20) as _, i}
      <div class="petal" style="--i: {i}"></div>
    {/each}
  </div>

  <div class="profile">
    <div class="avatar-container">
      <!-- Profile picture without link -->
      <img src={profilepic} alt="Queenayer Avatar" class="avatar" />
      <!-- Invisible proxy circle for clickable area -->
      <a href={profilepicLink} target="_blank" rel="noopener noreferrer" class="avatar-proxy-link" aria-label="View profile picture">
        <div class="avatar-proxy"></div>
      </a>
      <!-- Avatar decoration -->
      <a href={avatardecoLink} target="_blank" rel="noopener noreferrer" class="avatar-deco-link" aria-label="View avatar decoration">
        <img src={avatardeco} alt="Avatar Decoration" class="avatar-deco" />
      </a>
    </div>
    <h1 class="username">Queenayer</h1>
    <p class="bio">{bio}</p>
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

  <footer class="footer">
    Made with ❤️ by <a href="/synertry" target="_blank" rel="noopener noreferrer">Synertry</a>
  </footer>

  <!-- Conillo mascot -->
  <div class="mascot-container">
    <a href={conilloLink} target="_blank" rel="noopener noreferrer" class="mascot-link" aria-label="View Conillo mascot">
      <img src={conillo} alt="Conillo mascot" class="mascot" />
    </a>
  </div>
</main>

