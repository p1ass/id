import { notFound } from 'next/navigation'

import { AuthorizeRequestSchema } from '../../../../lib/oauth2/types'

type PageProps = {
  // Workaround: If we remove ? from searchParams, we get compile error
  searchParams?: { [key: string]: string | string[] | undefined }
}

const AuthorizePage = async ({ searchParams }: PageProps) => {
  if (!searchParams) {
    return notFound()
  }

  const parsed = AuthorizeRequestSchema.safeParse(searchParams)
  if (!parsed.success) {
    // TODO: Error Handling
    console.log(parsed.error)
    throw parsed.error
  }
  const parameter = parsed.data

  return (
    <div>
      <h2>Consent Page</h2>
      <form action="/api/oauth2/authorize" method="post">
        <div>
          <label htmlFor="client_id">client_id: </label>
          <input name="client_id" value={parameter.client_id} />
        </div>
        <div>
          <label htmlFor="redirect_uri">redirect_uri: </label>
          <input name="redirect_uri" value={parameter.redirect_uri} />
        </div>
        <div>
          <label htmlFor="scope">scope: </label>
          <input name="scope" value={parameter.scope} />
        </div>
        <div>
          <label htmlFor="state">state: </label>
          <input name="state" value={parameter.state} />
        </div>
        <div>
          <label htmlFor="response_type">response_type: </label>
          <input name="response_type" value={parameter.response_type} />
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
