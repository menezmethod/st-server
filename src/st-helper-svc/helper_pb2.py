# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: helper.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0chelper.proto\x12\x06helper\";\n\x10\x46inancialRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x16\n\x0e\x66inancial_data\x18\x02 \x01(\t\"(\n\x11\x46inancialResponse\x12\x13\n\x0bsuggestions\x18\x01 \x01(\t\"5\n\x11StockQuoteRequest\x12\x0e\n\x06ticker\x18\x01 \x01(\t\x12\x10\n\x08provider\x18\x02 \x01(\t\"\x82\x03\n\x12StockQuoteResponse\x12\x0e\n\x06symbol\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x10\n\x08\x65xchange\x18\x03 \x01(\t\x12\x11\n\tlastPrice\x18\x04 \x01(\x02\x12\x0c\n\x04open\x18\x05 \x01(\x02\x12\x0c\n\x04high\x18\x06 \x01(\x02\x12\x0b\n\x03low\x18\x07 \x01(\x02\x12\x0e\n\x06volume\x18\x08 \x01(\x03\x12\x11\n\tprevClose\x18\t \x01(\x02\x12\x0e\n\x06\x63hange\x18\n \x01(\x02\x12\x15\n\rchangePercent\x18\x0b \x01(\x02\x12\x10\n\x08yearHigh\x18\x0c \x01(\x02\x12\x0f\n\x07yearLow\x18\r \x01(\x02\x12\x11\n\tmarketCap\x18\x0e \x01(\x03\x12\x19\n\x11sharesOutstanding\x18\x0f \x01(\x03\x12\n\n\x02pe\x18\x10 \x01(\x02\x12\x1c\n\x14\x65\x61rningsAnnouncement\x18\x11 \x01(\t\x12\x0b\n\x03\x65ps\x18\x12 \x01(\x02\x12\x0e\n\x06sector\x18\x13 \x01(\t\x12\x10\n\x08industry\x18\x14 \x01(\t\x12\x0c\n\x04\x62\x65ta\x18\x15 \x01(\x02\"d\n\x1aHistoricalStockDataRequest\x12\x0e\n\x06symbol\x18\x01 \x01(\t\x12\x12\n\nstart_date\x18\x02 \x01(\t\x12\x10\n\x08\x65nd_date\x18\x03 \x01(\t\x12\x10\n\x08provider\x18\x04 \x01(\t\"\xc9\x01\n\x1bHistoricalStockDataResponse\x12\x0c\n\x04open\x18\x01 \x03(\x02\x12\x0c\n\x04high\x18\x02 \x03(\x02\x12\x0b\n\x03low\x18\x03 \x03(\x02\x12\r\n\x05\x63lose\x18\x04 \x03(\x02\x12\x0e\n\x06volume\x18\x05 \x03(\x03\x12\x0c\n\x04vwap\x18\x06 \x03(\x02\x12\x11\n\tadj_close\x18\x07 \x03(\x02\x12\x19\n\x11unadjusted_volume\x18\x08 \x03(\x02\x12\x0e\n\x06\x63hange\x18\t \x03(\x02\x12\x16\n\x0e\x63hange_percent\x18\n \x03(\x02\x32\x88\x02\n\x08STHelper\x12M\n\x14\x41nalyzeFinancialData\x12\x18.helper.FinancialRequest\x1a\x19.helper.FinancialResponse\"\x00\x12H\n\rGetStockQuote\x12\x19.helper.StockQuoteRequest\x1a\x1a.helper.StockQuoteResponse\"\x00\x12\x63\n\x16GetHistoricalStockData\x12\".helper.HistoricalStockDataRequest\x1a#.helper.HistoricalStockDataResponse\"\x00\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'helper_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_FINANCIALREQUEST']._serialized_start=24
  _globals['_FINANCIALREQUEST']._serialized_end=83
  _globals['_FINANCIALRESPONSE']._serialized_start=85
  _globals['_FINANCIALRESPONSE']._serialized_end=125
  _globals['_STOCKQUOTEREQUEST']._serialized_start=127
  _globals['_STOCKQUOTEREQUEST']._serialized_end=180
  _globals['_STOCKQUOTERESPONSE']._serialized_start=183
  _globals['_STOCKQUOTERESPONSE']._serialized_end=569
  _globals['_HISTORICALSTOCKDATAREQUEST']._serialized_start=571
  _globals['_HISTORICALSTOCKDATAREQUEST']._serialized_end=671
  _globals['_HISTORICALSTOCKDATARESPONSE']._serialized_start=674
  _globals['_HISTORICALSTOCKDATARESPONSE']._serialized_end=875
  _globals['_STHELPER']._serialized_start=878
  _globals['_STHELPER']._serialized_end=1142
# @@protoc_insertion_point(module_scope)