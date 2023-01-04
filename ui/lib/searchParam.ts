export const convertToSearchParam = (parameter: {
  [s: string]: string | null | undefined
}): URLSearchParams => {
  return new URLSearchParams(
    Object.entries(parameter).filter((kv): kv is [string, string] => kv[1] != null)
  )
}
