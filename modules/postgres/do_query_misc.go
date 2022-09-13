// SPDX-License-Identifier: GPL-3.0-or-later

package postgres

import "strconv"

func (p *Postgres) doQuerySettingsMaxConnections() (int64, error) {
	q := querySettingsMaxConnections()

	var s string
	if err := p.doQueryRow(q, &s); err != nil {
		return 0, err
	}

	return strconv.ParseInt(s, 10, 64)
}

func (p *Postgres) doQueryServerVersion() (int, error) {
	q := queryServerVersion()

	var s string
	if err := p.doQueryRow(q, &s); err != nil {
		return 0, err
	}

	return strconv.Atoi(s)
}

func (p *Postgres) doQueryIsSuperUser() (bool, error) {
	q := queryIsSuperUser()

	var v bool
	if err := p.doQueryRow(q, &v); err != nil {
		return false, err
	}

	return v, nil
}

func (p *Postgres) doQueryPGIsInRecovery() (bool, error) {
	q := queryPGIsInRecovery()

	var v bool
	if err := p.doQueryRow(q, &v); err != nil {
		return false, err
	}

	return v, nil
}

func (p *Postgres) doQueryCurrentDB() (string, error) {
	q := queryCurrentDatabase()

	var s string
	if err := p.doQueryRow(q, &s); err != nil {
		return "", err
	}

	return s, nil
}

func (p *Postgres) doQueryQueryableDatabases() error {
	q := queryQueryableDatabaseList()

	var dbs []string
	err := p.doQuery(q, func(_, value string, _ bool) {
		if p.dbSr != nil && p.dbSr.MatchString(value) {
			dbs = append(dbs, value)
		}
	})
	if err != nil {
		return err
	}

	seen := make(map[string]bool, len(dbs))

	for _, dbname := range dbs {
		seen[dbname] = true

		conn, ok := p.dbConns[dbname]
		if !ok {
			conn = &dbConn{}
			p.dbConns[dbname] = conn
		}

		if conn.db != nil || conn.connErrors >= 3 {
			continue
		}

		var err error
		if conn.db, err = p.openSecondaryConnection(dbname); err != nil {
			p.Warning(err)
			conn.connErrors++
		}
	}

	for dbname, conn := range p.dbConns {
		if seen[dbname] {
			continue
		}
		delete(p.dbConns, dbname)
		if conn.db != nil {
			_ = conn.db.Close()
		}
	}

	return nil
}
