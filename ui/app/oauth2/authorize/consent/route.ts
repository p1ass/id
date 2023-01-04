import { AuthorizeRequest } from '../../../../lib/oauth2/types'
import { convertToSearchParam } from '../../../../lib/searchParam'

const path = '/oauth2/authorize/consent' as const

export const buildAuthorizeConsentUri = (query: AuthorizeRequest) => {
  const searchParam = convertToSearchParam(query)
  return `${path}?${searchParam.toString()}`
}
