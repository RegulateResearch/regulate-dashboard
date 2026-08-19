<script lang="ts">
	import { resolve } from '$app/paths';
	import AuthForm from '$lib/components/auth-form.svelte';
	import { Button } from '$lib/components/ui/button';
	import { Field, FieldDescription } from '$lib/components/ui/field/index.js';
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input/index';
	import * as Tooltip from '$lib/components/ui/tooltip';
	import Eye from '@lucide/svelte/icons/eye';
	import EyeOff from '@lucide/svelte/icons/eye-off';
	import { untrack } from 'svelte';
	import { type Infer, superForm, type SuperValidated } from 'sveltekit-superforms';
	import { zod4Client } from 'sveltekit-superforms/adapters';
	import { formSchema, type FormSchema } from './schema';

	let {
		data
	}: {
		data: { form: SuperValidated<Infer<FormSchema>> };
	} = $props();

	let showPassword = $state(false);
	let showPasswordConfirmation = $state(false);

	const form = superForm(
		untrack(() => data.form),
		{
			validators: zod4Client(formSchema)
		}
	);

	const { form: formData, enhance } = form;
</script>

<AuthForm formTitle="Daftar" {formBody} {switchFormBtn} />

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
		<Form.Field {form} name="username">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Username</Form.Label>
					<Input
						{...props}
						id="username"
						type="text"
						placeholder="Username"
						required
						bind:value={$formData.username}
					/>
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="displayName">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Nama Lengkap</Form.Label>
					<Input
						{...props}
						id="displayName"
						type="text"
						placeholder="Nama Lengkap"
						required
						bind:value={$formData.displayName}
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
		<Form.Field {form} name="passwordConfirmation">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Ulangi Kata Sandi</Form.Label>
					<div class="flex items-center gap-2">
						<Input
							{...props}
							id="passwordConfirmation"
							type={showPasswordConfirmation ? 'text' : 'password'}
							placeholder="&#9679;&#9679;&#9679;&#9679;&#9679;"
							required
							bind:value={$formData.passwordConfirmation}
						/>
						<Tooltip.Root>
							<Tooltip.Trigger>
								<Button
									variant="outline"
									size="icon"
									type="button"
									onclick={() => (showPasswordConfirmation = !showPasswordConfirmation)}
								>
									{#if showPasswordConfirmation}
										<EyeOff />
									{:else}
										<Eye />
									{/if}
								</Button>
							</Tooltip.Trigger>
							<Tooltip.Content collisionPadding={16}>
								<p>{showPasswordConfirmation ? 'Sembunyikan' : 'Tampilkan'} kata sandi</p>
							</Tooltip.Content>
						</Tooltip.Root>
					</div>
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Button type="submit" class="w-full">Daftar</Form.Button>
	</form>
{/snippet}

{#snippet switchFormBtn()}
	<Field class="p-2">
		<FieldDescription class="text-center">
			Sudah punya akun? <a href={resolve('/login')}>Masuk</a>
		</FieldDescription>
	</Field>
{/snippet}
