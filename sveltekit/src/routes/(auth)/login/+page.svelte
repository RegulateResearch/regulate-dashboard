<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input/index';
	import { formSchema, type FormSchema } from './schema';
	import { zod4Client } from 'sveltekit-superforms/adapters';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import AuthForm from '../authForm.svelte';
	import type { ActionData } from './$types';
	import Button from '$lib/components/ui/button/button.svelte';
	import Eye from '@lucide/svelte/icons/eye';
	import EyeOff from '@lucide/svelte/icons/eye-off';
	import * as Tooltip from '$lib/components/ui/tooltip';
	import { resolve } from '$app/paths';

	type FormValidationData = {
		data: { form: SuperValidated<Infer<FormSchema>> };
		form: ActionData;
	};

	let { data, form: actionData }: FormValidationData = $props();

	let showPassword = $state(false);

	// svelte-ignore state_referenced_locally
	const form = superForm(data.form, {
		validators: zod4Client(formSchema)
	});

	const { form: formData, enhance } = form;
</script>

<AuthForm formTitle="Masuk" showLogo showSSOButton {formBody} {switchFormBtn} />

{#snippet formBody()}
	<form method="POST" use:enhance class="flex flex-col gap-6">
		<Form.Field {form} name="email">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Email</Form.Label>
					<Input
						{...props}
						id="email"
						type="email"
						placeholder="email@mail.com"
						required
						bind:value={$formData.email}
					/>
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="password">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Kata Sandi</Form.Label>
					<div class="flex items-center gap-2">
						<Input
							{...props}
							id="password"
							type={showPassword ? 'text' : 'password'}
							placeholder="&#9679;&#9679;&#9679;&#9679;&#9679;"
							required
							bind:value={$formData.password}
						/>
						<Tooltip.Root>
							<Tooltip.Trigger>
								<Button
									variant="outline"
									size="icon"
									type="button"
									onclick={() => (showPassword = !showPassword)}
								>
									{#if showPassword}
										<EyeOff />
									{:else}
										<Eye />
									{/if}
								</Button>
							</Tooltip.Trigger>
							<Tooltip.Content collisionPadding={16}>
								<p>{showPassword ? 'Sembunyikan' : 'Tampilkan'} kata sandi</p>
							</Tooltip.Content>
						</Tooltip.Root>
					</div>
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Button type="submit" class="w-full">Masuk</Form.Button>
		{#if actionData?.message === 'failed'}
			<p class="text-center text-red-500">
				Email atau kata sandi anda tidak sesuai.<br /> Mohon periksa kembali email dan kata sandi Anda.
			</p>
		{/if}
	</form>
{/snippet}

{#snippet switchFormBtn()}
	<span>Belum punya akun? <a href={resolve("/register")} class="text-yellow-500">Daftar</a></span>
{/snippet}
