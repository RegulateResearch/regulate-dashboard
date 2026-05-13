<script lang="ts">
	import * as Tabs from '$lib/components/ui/tabs/index.js';
	import * as Card from '$lib/components/ui/card/index.js';
</script>

<Tabs.Root value="srl-profile" class="h-full w-full">
  <Tabs.List>
    <Tabs.Trigger value="srl-profile">Profil SRL</Tabs.Trigger>
    <Tabs.Trigger value="learning-objectives">Tujuan Pembelajaran</Tabs.Trigger>
    <Tabs.Trigger value="gradebook">Gradebook</Tabs.Trigger>
    <Tabs.Trigger value="activities">Aktivitas</Tabs.Trigger>
  </Tabs.List>
  <Tabs.Content value="srl-profile">
    <Card.Root class="h-full w-full overflow-auto">
      <Card.Header>
        <Card.Title>Profil SRL</Card.Title>
        <Card.Description>Update your SRL profile here.</Card.Description>
      </Card.Header>
      <Card.Content>
        <p>Change your SRL profile here.</p>
      </Card.Content>
    </Card.Root>
  </Tabs.Content>
  <Tabs.Content value="learning-objectives">
    <Card.Root class="h-full w-full overflow-auto">
      <Card.Header>
        <Card.Title>Tujuan Pembelajaran</Card.Title>
        <Card.Description>Update your learning objectives here.</Card.Description>
      </Card.Header>
      <Card.Content>
        <p>Change your learning objectives here.</p>
      </Card.Content>
    </Card.Root>
  </Tabs.Content>
  <Tabs.Content value="gradebook">
    <Card.Root class="h-full w-full overflow-auto">
      <Card.Header>
        <Card.Title>Gradebook</Card.Title>
        <Card.Description>View your grades here.</Card.Description>
      </Card.Header>
      <Card.Content>
        <p>View your grades here.</p>
      </Card.Content>
    </Card.Root>
  </Tabs.Content>
  <Tabs.Content value="activities">
    <Card.Root class="h-full w-full overflow-auto">
      <Card.Header>
        <Card.Title>Aktivitas</Card.Title>
        <Card.Description>View your activities here.</Card.Description>
      </Card.Header>
      <Card.Content>
        <p>View your activities here.</p>
      </Card.Content>
    </Card.Root>
  </Tabs.Content>
</Tabs.Root>