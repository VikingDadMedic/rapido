import supportedLocales from '../../../../lang/languages.json'


// getSupportedLocales() imports lang/lang.json instead of using the API to call it, returns an array with the values of lang.json annotated with code , name.
export function getSupportedLocales () {
  let annotatedLocales = []

  for (const code of Object.keys(supportedLocales)) {
    annotatedLocales.push({
      code,
      name: supportedLocales[code]
    })
  }

  return annotatedLocales
}

export function supportedLocalesInclude (locale) {
  return Object.keys(supportedLocales).includes(locale)
}
