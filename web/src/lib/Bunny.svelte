<!--
  -           QueenayerLinkTree
  -     Copyright (c) Synertry 2025.
  - Distributed under the Boost Software License, Version 1.0.
  -     (See accompanying file LICENSE or copy at
  -           https://www.boost.org/LICENSE_1_0.txt)
  -->

<script lang="ts">
  import { onMount } from 'svelte';

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

<style>
  /* Bunny position styles */
  .bunny {
    position: fixed;
    width: 40px;
    height: 40px;
    z-index: 3;
    pointer-events: none;
    transition: transform 0.2s ease-out;
    /* Add shadow for the bunny */
    filter: drop-shadow(0 3px 3px rgba(0, 0, 0, 0.3));
    left: var(--x);
    bottom: var(--y, 10%); /* Position at the bottom with default 10% */
  }

  .bunny.jumping {
    animation: hop 1s ease-in-out;
  }

  .bunny-head {
    position: absolute;
    width: 20px;
    height: 20px;
    background-color: white;
    border-radius: 50%;
    top: 10px;
    left: 10px;
    z-index: 2;
  }

  .bunny-ears {
    position: absolute;
    top: 0;
    left: 8px;
    z-index: 1;
  }

  .bunny-ears::before, .bunny-ears::after {
    content: '';
    position: absolute;
    width: 6px;
    height: 15px;
    background-color: white;
    border-radius: 6px 6px 0 0;
  }

  .bunny-ears::before {
    left: 0;
    transform: rotate(-10deg);
  }

  .bunny-ears::after {
    left: 10px;
    transform: rotate(10deg);
  }

  .bunny-body {
    position: absolute;
    width: 25px;
    height: 15px;
    background-color: white;
    border-radius: 12px;
    top: 25px;
    left: 8px;
  }

  .bunny-tail {
    position: absolute;
    width: 8px;
    height: 8px;
    background-color: white;
    border-radius: 50%;
    top: 27px;
    left: 2px;
  }

  /* Faster, more energetic hopping animation */
  @keyframes hop {
    0% {
      transform: translateY(0) scaleY(1);
    }
    10% {
      transform: translateY(0) scaleY(0.85); /* Compress before jump - faster */
    }
    35% {
      transform: translateY(-30px) scaleY(1.15); /* Stretch at peak - higher jump */
    }
    60% {
      transform: translateY(-10px) scaleY(1); /* Start falling - faster descent */
    }
    80% {
      transform: translateY(0) scaleY(0.9); /* Compress on landing */
    }
    100% {
      transform: translateY(0) scaleY(1); /* Return to normal */
    }
  }
</style>