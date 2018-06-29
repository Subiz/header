// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package auth

@SerialVersionUID(0L)
final case class UserAuth(
    userId: _root_.scala.Predef.String = "",
    method: scala.Option[auth.Method] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[UserAuth] with scalapb.lenses.Updatable[UserAuth] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (userId != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(1, userId) }
      if (method.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(method.get.serializedSize) + method.get.serializedSize }
      __size
    }
    final override def serializedSize: _root_.scala.Int = {
      var read = __serializedSizeCachedValue
      if (read == 0) {
        read = __computeSerializedValue()
        __serializedSizeCachedValue = read
      }
      read
    }
    def writeTo(`_output__`: _root_.com.google.protobuf.CodedOutputStream): _root_.scala.Unit = {
      {
        val __v = userId
        if (__v != "") {
          _output__.writeString(1, __v)
        }
      };
      method.foreach { __v =>
        _output__.writeTag(2, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): auth.UserAuth = {
      var __userId = this.userId
      var __method = this.method
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __userId = _input__.readString()
          case 18 =>
            __method = Option(_root_.scalapb.LiteParser.readMessage(_input__, __method.getOrElse(auth.Method.defaultInstance)))
          case tag => _input__.skipField(tag)
        }
      }
      auth.UserAuth(
          userId = __userId,
          method = __method
      )
    }
    def withUserId(__v: _root_.scala.Predef.String): UserAuth = copy(userId = __v)
    def getMethod: auth.Method = method.getOrElse(auth.Method.defaultInstance)
    def clearMethod: UserAuth = copy(method = None)
    def withMethod(__v: auth.Method): UserAuth = copy(method = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => {
          val __t = userId
          if (__t != "") __t else null
        }
        case 2 => method.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => _root_.scalapb.descriptors.PString(userId)
        case 2 => method.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = auth.UserAuth
}

object UserAuth extends scalapb.GeneratedMessageCompanion[auth.UserAuth] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[auth.UserAuth] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): auth.UserAuth = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    auth.UserAuth(
      __fieldsMap.getOrElse(__fields.get(0), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[auth.Method]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[auth.UserAuth] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      auth.UserAuth(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[auth.Method]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = AuthProto.javaDescriptor.getMessageTypes.get(3)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = AuthProto.scalaDescriptor.messages(3)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 2 => __out = auth.Method
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = auth.UserAuth(
  )
  implicit class UserAuthLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, auth.UserAuth]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, auth.UserAuth](_l) {
    def userId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.userId)((c_, f_) => c_.copy(userId = f_))
    def method: _root_.scalapb.lenses.Lens[UpperPB, auth.Method] = field(_.getMethod)((c_, f_) => c_.copy(method = Option(f_)))
    def optionalMethod: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[auth.Method]] = field(_.method)((c_, f_) => c_.copy(method = f_))
  }
  final val USER_ID_FIELD_NUMBER = 1
  final val METHOD_FIELD_NUMBER = 2
}