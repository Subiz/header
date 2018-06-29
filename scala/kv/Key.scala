// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package kv

@SerialVersionUID(0L)
final case class Key(
    partition: _root_.scala.Predef.String = "",
    key: _root_.scala.Predef.String = ""
    ) extends scalapb.GeneratedMessage with scalapb.Message[Key] with scalapb.lenses.Updatable[Key] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (partition != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, partition) }
      if (key != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, key) }
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
        val __v = partition
        if (__v != "") {
          _output__.writeString(2, __v)
        }
      };
      {
        val __v = key
        if (__v != "") {
          _output__.writeString(3, __v)
        }
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): kv.Key = {
      var __partition = this.partition
      var __key = this.key
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 18 =>
            __partition = _input__.readString()
          case 26 =>
            __key = _input__.readString()
          case tag => _input__.skipField(tag)
        }
      }
      kv.Key(
          partition = __partition,
          key = __key
      )
    }
    def withPartition(__v: _root_.scala.Predef.String): Key = copy(partition = __v)
    def withKey(__v: _root_.scala.Predef.String): Key = copy(key = __v)
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 2 => {
          val __t = partition
          if (__t != "") __t else null
        }
        case 3 => {
          val __t = key
          if (__t != "") __t else null
        }
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 2 => _root_.scalapb.descriptors.PString(partition)
        case 3 => _root_.scalapb.descriptors.PString(key)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = kv.Key
}

object Key extends scalapb.GeneratedMessageCompanion[kv.Key] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[kv.Key] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): kv.Key = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    kv.Key(
      __fieldsMap.getOrElse(__fields.get(0), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(1), "").asInstanceOf[_root_.scala.Predef.String]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[kv.Key] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      kv.Key(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).map(_.as[_root_.scala.Predef.String]).getOrElse("")
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = KvProto.javaDescriptor.getMessageTypes.get(0)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = KvProto.scalaDescriptor.messages(0)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = kv.Key(
  )
  implicit class KeyLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, kv.Key]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, kv.Key](_l) {
    def partition: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.partition)((c_, f_) => c_.copy(partition = f_))
    def key: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.key)((c_, f_) => c_.copy(key = f_))
  }
  final val PARTITION_FIELD_NUMBER = 2
  final val KEY_FIELD_NUMBER = 3
}