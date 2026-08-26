<script lang="ts">
	import { onMount } from 'svelte';
	import { quintOut } from 'svelte/easing';
	import type { ClassValue } from 'svelte/elements';
	import { draw, fade } from 'svelte/transition';

	type Props = {
		animated?: boolean;
		loop?: boolean;
		withTitle?: boolean;
		active?: boolean;
		class?: ClassValue | null;
	};

	let {
		animated = false,
		loop = false,
		withTitle = false,
		active = true,
		class: className = null
	}: Props = $props();

	let visible = $derived(!animated);
	let fillClass: ClassValue = $derived(active ? 'fill-yellow-300' : 'fill-neutral-300');

	onMount(() => {
		if (animated && !loop) visible = true;
		if (animated && loop) {
			const interval = setInterval(() => {
				visible = !visible;
			}, 2000);

			return () => clearInterval(interval);
		}
	});
</script>

{#if visible}
	<svg
		xmlns="http://www.w3.org/2000/svg"
		width={withTitle ? 777 : 16}
		height={withTitle ? 195 : 16}
		viewBox={withTitle ? '0 0 777 195' : '0 30 120 135'}
		class={className}
		fill="none"
	>
		<path
			d="M76.7014 123.38C88.9565 119.876 92.9105 116.298 98.6963 109.606C104.482 102.914 108.124 94.6387 109.149 85.8522C110.174 77.0657 108.536 68.1739 104.446 60.3299C100.356 52.4859 94.0044 46.0516 86.2135 41.8616C78.4227 37.6716 69.5526 35.9191 60.7537 36.8314C51.9548 37.7437 43.633 41.2788 36.8676 46.9781C30.1021 52.6774 27.5359 56.8812 25.4366 64.3515C23.0433 72.8677 25.7087 86.0023 34.4333 89.8922C55.9063 97.8011 48.1392 49.5558 24.8067 74.2434C21.2345 78.023 17.2976 83.512 13.0449 91.1736C2.81367 109.606 -0.92041 123.961 20.5795 133.5C43.4937 143.666 100.15 118.944 74.1156 152.114C48.0813 185.284 37.2693 69.0066 18.2469 101.903C-0.775494 134.799 64.4463 126.885 76.7014 123.38Z"
			stroke="#292524"
			stroke-width="6"
			in:draw={{ duration: 500, delay: 0, easing: quintOut }}
			out:draw={{ duration: 500, delay: 200, easing: quintOut }}
		/>
		<path
			d="M40.2429 73.0677C42.482 75.6996 42.5731 84.4855 40.0608 85.5668C35.3673 87.5872 30.288 80.4087 30.2542 76.922C30.238 75.2421 38.1969 70.6628 40.2429 73.0677Z"
			class={fillClass}
			in:fade={{ duration: 250, delay: 0, easing: quintOut }}
			out:fade={{ duration: 250, delay: 400, easing: quintOut }}
		/>
		<path
			d="M73.8202 141.584C74.9181 144.86 68.9306 152.703 66.1964 152.77C61.0881 152.898 55.1828 143.637 56.45 140.388C57.0605 138.823 72.817 138.59 73.8202 141.584Z"
			class={fillClass}
			in:fade={{ duration: 250, delay: 100, easing: quintOut }}
			out:fade={{ duration: 250, delay: 300, easing: quintOut }}
		/>
		<path
			d="M20.3843 113.088C20.3024 109.633 22.7969 103.321 24.7767 103.494C29.695 104.88 34.255 119.166 33.4131 120.622C32.5721 122.076 20.4592 116.245 20.3843 113.088Z"
			class={fillClass}
			in:fade={{ duration: 250, delay: 200, easing: quintOut }}
			out:fade={{ duration: 250, delay: 200, easing: quintOut }}
		/>
		<path
			d="M98.6415 93.7218C80.4832 132.116 42.6194 120.608 44.8241 111.496C70.1768 59.5022 38.5591 66.866 37.0151 59.0195C35.471 51.173 62.8061 35.8352 84.0152 48.0996C99.9402 57.3083 106.472 77.1645 98.6415 93.7218Z"
			class={fillClass}
			in:fade={{ duration: 250, delay: 300, easing: quintOut }}
			out:fade={{ duration: 250, delay: 100, easing: quintOut }}
		/>
		{#if withTitle}
			<text
				x="143.633"
				y="149"
				fill="#292524"
				font-size="150"
				class="font-extralight"
				in:fade={{ duration: 500, delay: 500, easing: quintOut }}
				out:fade={{ duration: 500, delay: 500, easing: quintOut }}
			>
				regulate
			</text>
		{/if}
	</svg>
{/if}
