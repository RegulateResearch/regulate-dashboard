<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input/index';
	import { formSchema, type FormSchema } from './schema';
	import { zod4Client } from 'sveltekit-superforms/adapters';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import AuthForm from '../authForm.svelte';

	let {
		data
	}: {
		data: { form: SuperValidated<Infer<FormSchema>> };
	} = $props();

	// svelte-ignore state_referenced_locally
	const form = superForm(data.form, {
		validators: zod4Client(formSchema)
	});

	const { form: formData, enhance } = form;
</script>

<AuthForm formTitle="Daftar" showLogo showSSOButton {formBody} {switchFormBtn} />

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
					<Input
						{...props}
						id="password"
						type="password"
						placeholder="&#9679;&#9679;&#9679;&#9679;&#9679;"
						required
						bind:value={$formData.password}
					/>
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="passwordConfirmation">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Ulangi Kata Sandi</Form.Label>
					<Input
						{...props}
						id="password"
						type="password"
						placeholder="&#9679;&#9679;&#9679;&#9679;&#9679;"
						required
						bind:value={$formData.passwordConfirmation}
					/>
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Button type="submit" class="w-full">Daftar</Form.Button>
	</form>
{/snippet}

{#snippet switchFormBtn()}
	<span>Sudah punya akun? <a href="/login" class="text-yellow-500">Masuk</a></span>
{/snippet}
