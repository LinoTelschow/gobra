package viper.gobra.frontend.info.implementation.resolution

import viper.gobra.ast.frontend._
import viper.gobra.frontend.info.base.SymbolTable.isDefinedInScope
import viper.gobra.frontend.info.implementation.TypeInfoImpl

trait AmbiguityResolution { this: TypeInfoImpl =>

  def resolveConversionOrUnaryCall[T](n: PConversionOrUnaryCall)
                                     (conversion: (PUseLikeId, PExpression) => T)
                                     (unaryCall: (PUseLikeId, PExpression) => T): Option[T] =
    if (pointsToType(n.base))      Some(conversion(n.base, n.arg))
    else if (pointsToData(n.base)) Some(unaryCall(n.base, n.arg))
    else None
  /*
  def resolveSelectionOrMethodExpr[T](n: PSelectionOrMethodExpr)
                                     (selection: (PUseLikeId, PIdnUse) => T)
                                     (methodExp: (PUseLikeId, PIdnUse) => T): Option[T] =
    if (pointsToType(n.base))      Some(methodExp(n.base, n.id))
    else if (pointsToData(n.base)) Some(selection(n.base, n.id))
    else None
  */
  def resolveMPredOrMethExprOrRecvCall[T](n: PMPredOrMethRecvOrExprCall)
                                         (predOrMethCall: (PUseLikeId, PIdnUse, Vector[PExpression]) => T)
                                         (predOrMethExprCall: (PUseLikeId, PIdnUse, Vector[PExpression]) => T)
                                         : Option[T] =
    if (pointsToType(n.base))      Some(predOrMethCall(n.base, n.id, n.args))
    else if (pointsToData(n.base)) Some(predOrMethExprCall(n.base, n.id, n.args))
    else None

  def isDef[T](n: PIdnUnk): Boolean = !isDefinedInScope(sequentialDefenv.in(n), serialize(n))
}
