export const fetchFiles = async () => {
  const res = await fetch('http://localhost:3000/api/files', {
    method: 'GET',
    headers: {
      Accept: 'application/json',
    },
  })
  console.log(res)
  const body = await res.json()
  console.log(body)
}
