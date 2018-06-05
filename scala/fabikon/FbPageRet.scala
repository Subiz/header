// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package fabikon

@SerialVersionUID(0L)
final case class FbPageRet(
    data: _root_.scala.collection.Seq[fabikon.FbPageEntry] = _root_.scala.collection.Seq.empty
    ) extends scalapb.GeneratedMessage with scalapb.Message[FbPageRet] with scalapb.lenses.Updatable[FbPageRet] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      data.foreach(data => __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(data.serializedSize) + data.serializedSize)
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
      data.foreach { __v =>
        _output__.writeTag(2, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): fabikon.FbPageRet = {
      val __data = (_root_.scala.collection.immutable.Vector.newBuilder[fabikon.FbPageEntry] ++= this.data)
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 18 =>
            __data += _root_.scalapb.LiteParser.readMessage(_input__, fabikon.FbPageEntry.defaultInstance)
          case tag => _input__.skipField(tag)
        }
      }
      fabikon.FbPageRet(
          data = __data.result()
      )
    }
    def clearData = copy(data = _root_.scala.collection.Seq.empty)
    def addData(__vs: fabikon.FbPageEntry*): FbPageRet = addAllData(__vs)
    def addAllData(__vs: TraversableOnce[fabikon.FbPageEntry]): FbPageRet = copy(data = data ++ __vs)
    def withData(__v: _root_.scala.collection.Seq[fabikon.FbPageEntry]): FbPageRet = copy(data = __v)
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 2 => data
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 2 => _root_.scalapb.descriptors.PRepeated(data.map(_.toPMessage)(_root_.scala.collection.breakOut))
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = fabikon.FbPageRet
}

object FbPageRet extends scalapb.GeneratedMessageCompanion[fabikon.FbPageRet] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[fabikon.FbPageRet] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): fabikon.FbPageRet = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    fabikon.FbPageRet(
      __fieldsMap.getOrElse(__fields.get(0), Nil).asInstanceOf[_root_.scala.collection.Seq[fabikon.FbPageEntry]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[fabikon.FbPageRet] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      fabikon.FbPageRet(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).map(_.as[_root_.scala.collection.Seq[fabikon.FbPageEntry]]).getOrElse(_root_.scala.collection.Seq.empty)
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = FabikonProto.javaDescriptor.getMessageTypes.get(20)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = FabikonProto.scalaDescriptor.messages(20)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 2 => __out = fabikon.FbPageEntry
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = fabikon.FbPageRet(
  )
  implicit class FbPageRetLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, fabikon.FbPageRet]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, fabikon.FbPageRet](_l) {
    def data: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[fabikon.FbPageEntry]] = field(_.data)((c_, f_) => c_.copy(data = f_))
  }
  final val DATA_FIELD_NUMBER = 2
}
