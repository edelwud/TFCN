add_library(libcomasync STATIC libcomasync.h libcomasync.c)
target_include_directories (libcomasync PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})