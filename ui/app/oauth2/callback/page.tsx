const CallbackPage = ({
  searchParams
}: {
  searchParams?: { [key: string]: string | string[] | undefined }
}) => {
  if (searchParams) {
    searchParams
  }
  return (
    <div>
      <h1>Callback page</h1>
      <p>{JSON.stringify(searchParams, null, 2)}</p>
    </div>
  )
}

export default CallbackPage
