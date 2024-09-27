export const load = async ({}) => {
  console.log('a')
	const res = await fetch('https://example.com/')
	console.log(res.status)

	return {}
}
