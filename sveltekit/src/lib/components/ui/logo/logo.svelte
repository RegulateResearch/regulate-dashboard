<script lang="ts">
	import { draw, fade, type TransitionConfig } from 'svelte/transition';
	import { quintOut } from 'svelte/easing';
	import { onMount } from 'svelte';
	import type { ClassValue } from 'clsx';
	import { cn } from '$lib/utils';

	type Props = {
		animated?: boolean;
		loop?: boolean;
		active?: boolean;
		withTitle?: boolean;
		withSubTitle?: boolean;
		class?: ClassValue;
	};

	let {
		animated = false,
		loop = false,
		active = true,
		class: className = null,
		withTitle = false,
		withSubTitle = false
	}: Props = $props();

	let visible = $derived(!animated);

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

<div class={cn('@container', className)}>
	<div class="flex flex-row items-center justify-center gap-[3cqw]">
		<svg
			class="h-full w-auto origin-bottom-right -translate-y-[6cqw]"
			width="143"
			height="179"
			viewBox="0 0 143 179"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
		>
			{#if visible}
				<path
					class={active ? 'fill-yellow-300' : 'fill-neutral-300'}
					in:fade={{ duration: 500, delay: 400, easing: quintOut }}
					out:fade={{ duration: 500, delay: 0, easing: quintOut }}
					d="M60.2461 118.124C51.2264 115.325 53.5504 110.349 50.5266 102.184C46.6388 91.6872 44.3062 90.1322 44.3062 80.0239C44.3062 69.9157 50.5266 62.5289 66.4665 51.2543C82.4065 39.9797 90.5219 50.1164 108.066 56.3084C114.675 58.6411 131.393 97.5189 128.282 102.962C124.034 110.396 110.787 115.403 94.4586 122.79C78.1299 130.176 71.5207 121.623 60.2461 118.124Z"
				/>
				<path
					in:draw={{ duration: 500, delay: 0, easing: quintOut }}
					out:draw={{ duration: 500, delay: 0, easing: quintOut }}
					d="M123.849 174.352H49.9452C40.025 174.013 39.9009 162.98 48.2869 162.718H122.191C133.163 163.744 133.164 150.719 122.762 151.165H49.2873C38.6383 150.267 39.5038 137.943 50.2993 139.134C77.0667 138.557 125.189 107.352 124.816 130.879C124.553 147.433 33.2419 150.821 34.3436 87.1845C34.957 51.7559 66.0856 35.6183 88.3915 36.0634C110.698 36.5085 143.39 51.7717 128.564 101.429L120.948 89.1918L137.51 89.65C137.51 89.65 128.453 102.719 128.334 101.241"
					stroke="black"
					stroke-width="9.12756"
					stroke-linecap="round"
					stroke-linejoin="round"
				/>
				<path
					in:draw={{ duration: 500, delay: 200, easing: quintOut }}
					out:draw={{ duration: 500, delay: 400, easing: quintOut }}
					d="M18.7312 117.016L8.65527 123.156"
					stroke="black"
					stroke-width="7.81228"
					stroke-linecap="round"
					stroke-linejoin="round"
				/>
				<path
					in:draw={{ duration: 500, delay: 300, easing: quintOut }}
					out:draw={{ duration: 500, delay: 300, easing: quintOut }}
					d="M15.0233 86.8739L4.30615 84.7698"
					stroke="black"
					stroke-width="7.81228"
					stroke-linecap="round"
					stroke-linejoin="round"
				/>
				<path
					in:draw={{ duration: 500, delay: 400, easing: quintOut }}
					out:draw={{ duration: 500, delay: 200, easing: quintOut }}
					d="M24.0943 47.8244L16.3008 40.9871"
					stroke="black"
					stroke-width="7.81228"
					stroke-linecap="round"
					stroke-linejoin="round"
				/>
				<path
					in:draw={{ duration: 500, delay: 500, easing: quintOut }}
					out:draw={{ duration: 500, delay: 100, easing: quintOut }}
					d="M52.2234 21.9649L48.4204 13.0469"
					stroke="black"
					stroke-width="7.81228"
					stroke-linecap="round"
					stroke-linejoin="round"
				/>
				<path
					in:draw={{ duration: 500, delay: 600, easing: quintOut }}
					out:draw={{ duration: 500, delay: 0, easing: quintOut }}
					d="M85.4102 14.9752L85.7408 5.17578"
					stroke="black"
					stroke-width="7.81228"
					stroke-linecap="round"
					stroke-linejoin="round"
				/>
			{/if}
		</svg>
		{#if withTitle}
			<div class="flex h-full flex-col items-start justify-center">
				<h1
					class={cn(
						'font-cal-sans text-[15cqw]',
						withSubTitle ? 'underline underline-offset-4' : ''
					)}
				>
					reguLAte!
				</h1>
				{#if withSubTitle}
					<p class="flex flex-row items-center gap-[2.5cqw] font-cal-sans text-[4cqw]">
						<span>plan</span>
						<span>&#9679;</span>
						<span>do</span>
						<span>&#9679;</span>
						<span>monitor</span>
						<span>&#9679;</span>
						<span>evaluate</span>
					</p>
				{/if}
			</div>
		{/if}
	</div>
</div>
