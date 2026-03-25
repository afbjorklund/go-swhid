/*
** 2026-03-25
**
** The author disclaims copyright to this source code.  In place of
** a legal notice, here is a blessing:
**
**    May you do good and not evil.
**    May you find forgiveness for yourself and forgive others.
**    May you share freely, never taking more than you give.
**
******************************************************************************
**
** This SQLite extension implements functions for git.
*/
#include "sqlite3ext.h"
SQLITE_EXTENSION_INIT1
#include <assert.h>
#include <string.h>
#include <stdarg.h>

static const char* gittype[] = {
	"none",
	"commit", /* OBJ_COMMIT = 1 */
	"tree", /* OBJ_TREE = 2 */
	"blob", /* OBJ_BLOB = 3 */
	"tag", /* OBJ_TAG = 4 */
	"snapshot", /* 5 for future expansion */
};

static void gitObjectType(
  sqlite3_context *context,
  int argc,
  sqlite3_value **argv
){
  int eType = sqlite3_value_type(argv[0]);
  int nInt = sqlite3_value_int(argv[0]);

  assert( argc==1 );
  if( eType==SQLITE_INTEGER ){
    if (nInt >= 0 && nInt <= 5) {
      sqlite3_result_text(context, gittype[nInt], -1, SQLITE_STATIC);
    }
  } else if( eType == SQLITE_TEXT ){
    sqlite3_result_value(context, argv[0]);
  }
}

#ifdef _WIN32
__declspec(dllexport)
#endif
int sqlite3_git_init(
  sqlite3 *db,
  char **pzErrMsg,
  const sqlite3_api_routines *pApi
){
  int rc = SQLITE_OK;
  SQLITE_EXTENSION_INIT2(pApi);
  (void)pzErrMsg;  /* Unused parameter */
  rc = sqlite3_create_function(db, "git_object_type", 1,
                       SQLITE_UTF8 | SQLITE_INNOCUOUS | SQLITE_DETERMINISTIC,
                                0, gitObjectType, 0, 0);
  return rc;
}
