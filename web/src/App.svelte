<script lang="ts">
  import { onMount } from 'svelte';
  import Icon from '@iconify/svelte';

  // Use static image from Twitter
  let profilePic = "https://pbs.twimg.com/media/GjL4pgDa8AAnvgd?format=jpg&name=large";

  // No dark mode - using only pastel pink theme

  // Social media links with actual URLs
  const links = [
    { name: 'Discord', url: 'https://discord.com/users/402581006185660423', icon: 'simple-icons:discord' },
    { name: 'Steam', url: 'https://steamcommunity.com/profiles/76561199206914943', icon: 'simple-icons:steam' },
    { name: 'Twitter', url: 'https://x.com/QueenAyer', icon: 'simple-icons:x' },
    { name: 'Reddit', url: 'https://www.reddit.com/u/QueenAyer', icon: 'simple-icons:reddit' },
    { name: 'Youtube', url: 'https://www.youtube.com/@queenayer2671', icon: 'simple-icons:youtube' },
    { name: 'Twitch', url: 'https://www.twitch.tv/queenayer', icon: 'simple-icons:twitch' },
  ];

  // Bunny position state
  let bunnyX = 10;
  let bunnyY = 10; // Position at the bottom
  let bunnyDirection = 1;
  let verticalDirection = 1; // For zig-zag movement
  let jumping = false;
  let hopProgress = 0;
  let hopDirection = 0;

  // Animate bunny
  function animateBunny() {
    // Only move horizontally when not jumping
    if (!jumping) {
      // Move horizontally with zig-zag pattern
      bunnyX += bunnyDirection * 0.3;

      // Add vertical zig-zag movement
      bunnyY += verticalDirection * 0.1;

      // Change vertical direction randomly for zig-zag effect
      if (Math.random() < 0.1) {
        verticalDirection = -verticalDirection;
      }

      // Keep bunny within vertical bounds
      if (bunnyY > 15) {
        verticalDirection = -1;
      } else if (bunnyY < 5) {
        verticalDirection = 1;
      }

      // Change direction at edges
      if (bunnyX > 80) {
        bunnyDirection = -1;
      } else if (bunnyX < 10) {
        bunnyDirection = 1;
      }

      // More frequent hopping
      if (Math.random() < 0.08) { // Much more frequent jumps
        jumping = true;
        hopProgress = 0;
        hopDirection = bunnyDirection; // Remember direction at start of hop
      }
    } else {
      // Faster hopping motion
      hopProgress += 0.05; // Faster progress through the hop

      // Move forward in the hop direction
      bunnyX += hopDirection * 0.2;

      // Add slight vertical movement during hop for zig-zag
      bunnyY += verticalDirection * 0.05;

      // Complete the hop
      if (hopProgress >= 1) {
        jumping = false;
        // Chance to change direction after hop for more zig-zag
        if (Math.random() < 0.3) {
          bunnyDirection = -bunnyDirection;
        }
      }
    }
  }

  // Start bunny animation
  onMount(() => {
    const interval = setInterval(animateBunny, 200); // Faster animation
    return () => clearInterval(interval);
  });
</script>

<main>

  <!-- Petal Animation Container -->
  <div class="petals-container">
    {#each Array(20) as _, i}
      <div class="petal" style="--i: {i}"></div>
    {/each}
  </div>

  <!-- Bunny Animation -->
  <div 
    class="bunny" 
    style="--x: {bunnyX}%; --y: {bunnyY}%;" 
    class:jumping={jumping}
  >
    <div class="bunny-ears"></div>
    <div class="bunny-head"></div>
    <div class="bunny-body"></div>
    <div class="bunny-tail"></div>
  </div>

  <div class="profile">
    <img src={profilePic} alt="Queenayer Avatar" class="avatar" />
    <h1 class="username">@Queenayer</h1>
    <p class="bio">Polish Weeb ✨ Exploring Teyvat, battling Honkai, and dominating the Rift~ Kocham Hoyoverse i League of Legends! ♡</p>
  </div>

  <div class="links">
    {#each links as link}
      <a href={link.url} target="_blank" rel="noopener noreferrer" class="link-button">
        <div class="icon-container">
          <Icon icon={link.icon} width="1.5em" height="1.5em" />
        </div>
        <span>{link.name}</span>
      </a>
    {/each}
  </div>

  <footer class="footer">
    Made with ❤️ by <a href="http://synertry.me" target="_blank" rel="noopener noreferrer">Synertry</a>
  </footer>
</main>

<style>
  /* Only keep dynamic styles for bunny position */
  .bunny {
    left: var(--x);
    bottom: var(--y, 10%); /* Position at the bottom with default 10% */
  }
</style>
