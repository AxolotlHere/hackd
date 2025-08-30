<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { browser } from '$app/environment';
	import logo from '$lib/assets/logo.png';

	let canvas: HTMLCanvasElement;
	let raf: number | null = null;

	const goToAuth = () => {
		window.location.href = '/auth';
	};

	if (browser) {
		onMount(() => {
			const ctx = canvas.getContext('2d')!;
			const letters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'.split('');

			const fontSize = 16;
			let columns: number;
			let drops: number[] = [];

			function sizeCanvas() {
				const dpr = Math.max(1, window.devicePixelRatio || 1);
				const w = canvas.clientWidth;
				const h = canvas.clientHeight;

				canvas.width = Math.floor(w * dpr);
				canvas.height = Math.floor(h * dpr);

				ctx.setTransform(dpr, 0, 0, dpr, 0, 0);

				columns = Math.ceil(w / fontSize);
				drops = Array.from({ length: columns }, () => Math.random() * -50);
			}

			function draw() {
				const w = canvas.clientWidth;
				const h = canvas.clientHeight;

				ctx.fillStyle = 'rgba(0, 0, 0, 0.075)';
				ctx.fillRect(0, 0, w, h);

				ctx.fillStyle = 'rgba(0, 255, 0, 0.5)';
				ctx.font = `${fontSize}px monospace`;

				for (let i = 0; i < drops.length; i++) {
					const ch = letters[(Math.random() * letters.length) | 0];
					ctx.fillText(ch, i * fontSize, drops[i] * fontSize);

					if (drops[i] * fontSize > h && Math.random() > 0.995) drops[i] = 0;
					drops[i] += 1;
				}

				raf = requestAnimationFrame(draw);
			}

			sizeCanvas();
			raf = requestAnimationFrame(draw);
			window.addEventListener('resize', sizeCanvas);

			onDestroy(() => {
				if (raf !== null) cancelAnimationFrame(raf);
				window.removeEventListener('resize', sizeCanvas);
			});
		});
	}
</script>

<div
	class="min-h-screen text-gray-200 flex flex-col items-center justify-center relative overflow-hidden"
>
	<!-- Matrix canvas background -->
	<canvas bind:this={canvas} class="matrix"></canvas>

	<img src={logo} alt="logo" class="w-40 mb-6 drop-shadow-[0_0_15px_rgba(0,255,255,0.7)]" />

	<h1 class=" text-7xl font-extrabold tracking-widest mb-6" data-text="HACK'd">HACK'd</h1>

	<p class="font-mono text-lg text-green-400 mb-12">
		&gt; Crack the code_ <span class="animate-pulse">█</span>
	</p>

	<button
		on:click={goToAuth}
		class="px-8 py-3 rounded-xl border border-cyan-400 text-cyan-400 hover:bg-cyan-400/20 hover:shadow-[0_0_25px_#00ffff] transition-all font-mono tracking-widest"
	>
		Enter the System
	</button>

	<footer class="absolute bottom-6 text-green-500 font-mono text-sm">
		&gt;_ May the force be with you
	</footer>
</div>

<style>
	.matrix {
		position: fixed;
		inset: 0;
		z-index: -1;
		background: #000;
		pointer-events: none;
		display: block;
		width: 100%;
		height: 100%;
	}

	.glitch {
		position: relative;
		text-transform: uppercase;
		color: white;
	}

	.glitch::before,
	.glitch::after {
		content: attr(data-text);
		position: absolute;
		left: 0;
		width: 100%;
		overflow: hidden;
		clip: rect(0, 900px, 0, 0);
	}

	.glitch::before {
		animation: glitchTop 1s infinite linear alternate-reverse;
		color: #af0;
	}

	.glitch::after {
		animation: glitchBottom 1s infinite linear alternate-reverse;
		color: black;
	}

	@keyframes glitchTop {
		0% {
			clip: rect(0, 9999px, 0, 0);
			transform: translate(0);
		}
		20% {
			clip: rect(0, 9999px, 85px, 0);
			transform: translate(-3px, -3px);
		}
		40% {
			clip: rect(0, 9999px, 0, 0);
			transform: translate(3px, -1px);
		}
		60% {
			clip: rect(0, 9999px, 95px, 0);
			transform: translate(-2px, 2px);
		}
		80% {
			clip: rect(0, 9999px, 0, 0);
			transform: translate(2px, 1px);
		}
		100% {
			clip: rect(0, 9999px, 75px, 0);
			transform: translate(0);
		}
	}

	@keyframes glitchBottom {
		0% {
			clip: rect(0, 9999px, 0, 0);
			transform: translate(0);
		}
		20% {
			clip: rect(0, 9999px, 45px, 0);
			transform: translate(3px, 2px);
		}
		40% {
			clip: rect(0, 9999px, 0, 0);
			transform: translate(-2px, -2px);
		}
		60% {
			clip: rect(0, 9999px, 65px, 0);
			transform: translate(2px, 1px);
		}
		80% {
			clip: rect(0, 9999px, 0, 0);
			transform: translate(-1px, -1px);
		}
		100% {
			clip: rect(0, 9999px, 50px, 0);
			transform: translate(0);
		}
	}
</style>
