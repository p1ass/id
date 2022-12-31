import { notFound } from 'next/navigation'

type PageProps = {
  // Workaround: If we remove ? from searchParams, we get compile error
  searchParams?: {
    client_id?: string
    redirect_uri?: string
    scope?: string
    state?: string
    response_type?: string
  }
}

const AuthorizePage = async ({ searchParams }: PageProps) => {
  if (!searchParams) {
    return notFound()
  }
  console.log(searchParams)

  return (
    <div>
      <h2>Consent Page</h2>
      <form action="/api/oauth2/authorize" method="post">
        <div>
          <label htmlFor="client_id">client_id: </label>
          <input name="client_id" value={searchParams.client_id} />
        </div>
        <div>
          <label htmlFor="redirect_uri">redirect_uri: </label>
          <input name="redirect_uri" value={searchParams.redirect_uri} />
        </div>
        <div>
          <label htmlFor="scope">scope: </label>
          <input name="scope" value={searchParams.scope} />
        </div>
        <div>
          <label htmlFor="state">state: </label>
          <input name="state" value={searchParams.state} />
        </div>
        <div>
          <label htmlFor="response_type">response_type: </label>
          <input name="response_type" value={searchParams.response_type} />
        </div>
        <div>
          <label htmlFor="consent">consent: </label>
          <input name="consent" value="true" />
        </div>
        <div>
          <input type="submit" value="Consent" />
        </div>
      </form>
    </div>
  )
}

export default AuthorizePage
