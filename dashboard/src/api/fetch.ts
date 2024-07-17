"use server"
export default async function fetchBackend({
  path,
  method,
  data,
}: {
  path: string
  method: string
  data: any
}) {
  const BASE_URL = process.env.BACKEND_HOST
  const res = await fetch(BASE_URL + path, {
    method: method,
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
    cache: "no-store",
  })

  const result = await res.json()
  if (res.ok) {
    console.log(result)
    return result
  } else {
    throw result
  }
}
