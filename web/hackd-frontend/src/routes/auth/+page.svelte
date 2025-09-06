<script lang="ts">
	//frontend
	import { onMount, onDestroy } from 'svelte';
	import { browser } from '$app/environment';
	import Root from '$lib/components/ui/tabs/tabs.svelte';
	import List from '$lib/components/ui/tabs/tabs-list.svelte';
	import Trigger from '$lib/components/ui/tabs/tabs-trigger.svelte';
	import Content from '$lib/components/ui/tabs/tabs-content.svelte';
	import google from '$lib/assets/icons8-google.svg';
	import github from '$lib/assets/icons8-github-120.svg';
	let canvas: HTMLCanvasElement;
	let raf: number | null = null;

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
	//----------------------backend---------------------------------
	let usr_email: string = '';
	let usr_pass: string = '';
	let user = null;
	async function fetchUser() {
		try {
			const res = await fetch(
				`http://localhost:8080/v1/getUser?email=${encodeURIComponent(usr_email)}`
			);
			if (res.ok) {
				user = await res.json();
				console.log('Enters');
				console.log(user);
			}
		} catch (err) {
			console.log(err.message);
		}
	}
</script>

<div
	class="min-h-screen text-gray-200 flex flex-col items-center justify-center relative overflow-hidden"
>
	<canvas bind:this={canvas} class="matrix"></canvas>

	<div class="flex flex-col items-center self-center">
		<Root value="login" class="align-middle m-auto">
			<List
				class="text-amber-50 bg-black/10 rounded-md bg-clip-padding backdrop-filter backdrop-blur-sm "
			>
				<Trigger value="login" class="bg-transparent text-[#21de54]">LOGIN</Trigger>
				<Trigger value="signup" class="bg-transparent text-[#21de54]">SIGN UP</Trigger>
			</List>
			<Content value="login">
				<div
					class="bg-black/5 border-1 border-[#21de54] rounded-md bg-clip-padding backdrop-filter backdrop-blur-sm w-[400px] h-[400px] flex flex-col"
				>
					<p class="text-center text-5xl font-sans font-extrabold text-[#21de54] pt-10">LOGIN</p>
					<input
						bind:value={usr_email}
						type="text"
						placeholder="Enter your email"
						class="w-[350px] bg-black/10 border border-[#21de54] rounded mx-2 mt-8 px-3 py-2 text-[#21de54] placeholder-[#21de54]"
					/>
					<input
						bind:value={usr_pass}
						type="password"
						placeholder="Enter your password"
						class="w-[350px] bg-black/10 border border-[#21de54] rounded mx-2 mt-8 px-3 py-2 text-[#21de54] placeholder-[#21de54]"
					/>
					<button
						class="w-[100px] my-[25px] self-center bg-[#21de54] text-black font-bold py-2 rounded hover:bg-[#1ab143] transition"
						on:click={fetchUser}>Login</button
					>
					<div class="flex items-center w-full text-[#21de54]">
						<hr class="flex-grow border-[#21de54]/50" />
						<span class="px-2 text-sm">or</span>
						<hr class="flex-grow border-[#21de54]/50" />
					</div>
					<p class="text-center text-md font-sans font-extrabold text-[#21de54] pt-10">
						SIGN IN WITH
					</p>
					<div class="flex flex-row justify-center m-5">
						<button
							class=" w-fit flex items-center justify-center gap-2 border border-black hover:border-[#21de54] rounded-full mx-4 px-3 py-3 hover:bg-[#fff] transition bg-[#21de54]"
						>
							<img src={google} alt="Google" class="w-6 h-6" />
						</button>

						<button
							class=" w-fit flex items-center justify-center gap-2 border border-black hover:border-[#21de54] rounded-full mx-4 px-3 py-3 hover:bg-[#fff] transition bg-[#21de54]"
						>
							<img src={github} alt="GitHub" class="w-6 h-6" />
						</button>
					</div>
				</div>
			</Content>
			<Content value="signup">
				<div
					class="w-[400px] border-1 border-[#21de54] h-[400px] flex flex-col rounded-lg overflow-scroll"
				>
					<p class="text-center text-5xl font-sans font-extrabold text-[#21de54] pt-10">SIGN UP</p>
					<input
						type="text"
						placeholder="Enter your username"
						class="w-[350px] bg-black/10 border border-[#21de54] rounded mx-2 mt-8 px-3 py-2 text-[#21de54] placeholder-[#21de54]"
					/>
					<textarea
						placeholder="Give a short description about yourself."
						rows="6"
						cols="20"
						class="text-start bg-black/10 border border-[#21de54] rounded mx-2 mt-8 px-3 py-2 text-[#21de54] placeholder-[#21de54]"
					></textarea>
					<input
						type="text"
						placeholder="Enter your email"
						class="w-[350px] bg-black/10 border border-[#21de54] rounded mx-2 mt-8 px-3 py-2 text-[#21de54] placeholder-[#21de54]"
					/>
					<input
						type="password"
						placeholder="Enter your password"
						class="w-[350px] bg-black/10 border border-[#21de54] rounded mx-2 mt-8 px-3 py-2 text-[#21de54] placeholder-[#21de54]"
					/>
					<button
						class="w-[100px] my-[25px] self-center bg-[#21de54] text-black font-bold py-2 rounded hover:bg-[#1ab143] transition"
						>SIGN UP</button
					>
					<div class="flex items-center w-full text-[#21de54]">
						<hr class="flex-grow border-[#21de54]/50" />
						<span class="px-2 text-sm">or</span>
						<hr class="flex-grow border-[#21de54]/50" />
					</div>
					<p class="text-center text-md font-sans font-extrabold text-[#21de54] pt-10">
						SIGN IN WITH
					</p>
					<div class="flex flex-row justify-center m-5">
						<button
							class=" w-fit flex items-center justify-center gap-2 border border-black hover:border-[#21de54] rounded-full mx-4 px-3 py-3 hover:bg-[#fff] transition bg-[#21de54]"
						>
							<img src={google} alt="Google" class="w-6 h-6" />
						</button>

						<button
							class=" w-fit flex items-center justify-center gap-2 border border-black hover:border-[#21de54] rounded-full mx-4 px-3 py-3 hover:bg-[#fff] transition bg-[#21de54]"
						>
							<img src={github} alt="GitHub" class="w-6 h-6" />
						</button>
					</div>
				</div>
			</Content>
		</Root>
	</div>
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
</style>
