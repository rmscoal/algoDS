/* eslint-disable new-cap */
export default class IsInstance {
  public static Solver(obj: any, classFunction: any): boolean {
    if (obj == null || typeof classFunction !== 'function') return false;

    let ctor: any = obj.constructor;

    while (ctor) {
      if (ctor === classFunction) {
        return true;
      }

      ctor = Object.getPrototypeOf(ctor.prototype)?.constructor;
    }

    return false;
  }
}
