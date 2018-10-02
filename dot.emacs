;; .emacs

(custom-set-variables
 ;; uncomment to always end a file with a newline
 ;'(require-final-newline t)
 ;; uncomment to disable loading of "default.el" at startup
 ;'(inhibit-default-init t)
 ;; default to unified diffs
 '(diff-switches "-u"))

;;; uncomment for CJK utf-8 support for non-Asian users
;; (require 'un-define)

;; emacs でgolang編集環境を作る
;; https://qiita.com/bussorenre/items/3e80e59d517db8aebf46

(setq package-archives '(("gnu" . "http://elpa.gnu.org/packages/")
                         ("marmalade" . "http://marmalade-repo.org/packages/")
                         ("melpa" . "http://melpa.milkbox.net/packages/")))

;; Goのパスを通す
(add-to-list 'exec-path (expand-file-name "/usr/local/go/bin/"))
(add-to-list 'load-path "/home/snaga/.emacs.d/elpa/go-mode-20180327.1530")
(add-to-list 'load-path "/home/snaga/.emacs.d/elpa/company-go-20170825.1643")
(add-to-list 'load-path "/home/snaga/.emacs.d/elpa/company-20180913.2311")

;; go get で入れたツールのパスを通す
(add-to-list 'exec-path (expand-file-name "/home/snaga/.go/bin"))

;; 必要なパッケージのロード
(require 'go-mode)
(require 'company-go)

;; 諸々の有効化、設定
(add-hook 'go-mode-hook 'company-mode)
(add-hook 'go-mode-hook 'flycheck-mode)
(add-hook 'go-mode-hook (lambda()
;;           (add-hook 'before-save-hook' 'gofmt-before-save)
           (local-set-key (kbd "M-.") 'godef-jump)
           (set (make-local-variable 'company-backends) '(company-go))
           (company-mode)
           (setq indent-tabs-mode nil)    ; タブを利用
           (setq c-basic-offset 4)        ; tabサイズを4にする
           (setq tab-width 4)))
