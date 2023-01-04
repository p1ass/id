import { AuthorizeRequest } from '../../../lib/oauth2/types'
import { convertToSearchParam } from '../../../lib/searchParam'

const path = '/api/oauth2/authorize' as const

export const buildAuthorizePath = () => {
  return `${path}`
}

export const buildAuthorizeUri = (query: AuthorizeRequest) => {
  const searchParam = convertToSearchParam(query)
  return `${path}?${searchParam.toString()}`
}
