export class GlobalFunctions {
  public static generateParams(filter: any): string {
    return Object.keys(filter).map(key => key + '=' + filter[key]).join('&');;
  }
}
