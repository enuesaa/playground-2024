<script lang="ts">
	let content = ""
</script>

<section>
	{content}
	<!-- <textarea bind:value={content} /> -->
	<div contenteditable="true" style="width: 300px; height: 200px; border: 1px solid black;">
		Edit this text and<span style="font-weight: 800;">change its style!</span>
	  </div>
</section>

<style lang="postcss">
	textarea {
		visibility: hidden;
	}
</style>
